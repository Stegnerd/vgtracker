package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vgtracker/backend/internal/config"
	"vgtracker/backend/internal/gamedetails"
	"vgtracker/backend/internal/igdb"
	"vgtracker/backend/internal/utils"

	"github.com/pkg/errors"
)

type SteamInternalMethods interface {
	GetOwnedGames() (*OwnedGamesResponse, error)
	SyncOwnedGames() error
}

type SteamInternalHandler struct {
	config             *config.Config
	igdbClient         igdb.IgdbMethods
	gameDetailInternal gamedetails.GameDetailInternalMethods
}

func NewSteamInternal(
	config *config.Config,
	igdbClient *igdb.IgdbMethods,
	gameDetailInternal *gamedetails.GameDetailInternalMethods,
) SteamInternalMethods {
	return &SteamInternalHandler{
		config:             config,
		igdbClient:         *igdbClient,
		gameDetailInternal: *gameDetailInternal,
	}
}

type OwnedGamesResponse struct {
	Response struct {
		GameCount int        `json:"game_count"`
		Games     []GameInfo `json:"games"`
	} `json:"response"`
}

type GameInfo struct {
	AppID           int    `json:"appid"`
	Name            string `json:"name"`
	PlaytimeForever int    `json:"playtime_forever"`
	ImgIconURL      string `json:"img_icon_url"`
	PlaytimeWindows int    `json:"playtime_windows_forever"`
	PlaytimeMac     int    `json:"playtime_mac_forever"`
	PlaytimeLinux   int    `json:"playtime_linux_forever"`
}

func (s *SteamInternalHandler) GetOwnedGames() (*OwnedGamesResponse, error) {
	if s.config.Steam.SteamID == "" || s.config.Steam.APIKey == "" {
		return nil, errors.New("steam id or api key not set")
	}

	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json&include_appinfo=true&include_played_free_games=true",
		s.config.Steam.APIKey,
		s.config.Steam.SteamID,
	)

	// TODO: NEED TO ALSO MAKE SURE TO USE SEARCH BY URL AND MARRY IT UP TO IGDB AND SEARCH BY STEAM APPID TO IGDB STEAM URL

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get steam owned games")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("failed to get steam owned games: %s", resp.Status)
	}

	var gamesResponse OwnedGamesResponse
	if err := json.NewDecoder(resp.Body).Decode(&gamesResponse); err != nil {
		return nil, errors.WithMessage(err, "failed to decode steam games response")
	}

	return &gamesResponse, nil
}

// SyncOwnedGames
func (s *SteamInternalHandler) SyncOwnedGames() error {
	// 1. Get all owned games from steam
	steamGames, err := s.GetOwnedGames()
	if err != nil {
		return errors.WithMessage(err, "failed to get steam owned games")
	}

	// 2. run through each game getting the steam app id
	for _, game := range steamGames.Response.Games {
		appID := game.AppID

		// 3. check igdb via website url
		igdbGame, err := s.igdbClient.SearchGameBySteamAppID(appID)
		if err != nil {
			return errors.WithMessage(err, "failed to get igdb game via steam app id")
		}

		// 4. if igbd returns check against our game detail list
		if len(igdbGame.Items) == 1 {
			igdbGame := igdbGame.Items[0]
			gameDetail, err := s.gameDetailInternal.GetGameDetailRecord(igdbGame.GameID)
			if err != nil {
				return errors.WithMessage(err, "failed to get game detail record")
			}

			var id *int
			if gameDetail.ID != 0 {
				id = &gameDetail.ID
			}

			_, err = s.gameDetailInternal.UpsertGameDetailRecord(gamedetails.UpsertGameDetailInput{
				ID:           id,
				IGDBID:       &igdbGame.GameID,
				Name:         &igdbGame.Name,
				ThumbnailURL: &igdbGame.ThumbnailURL,
				CoverURL:     &igdbGame.CoverURL,
				IsOwned:      utils.Ptr(true),
				IsBeaten:     utils.Ptr(false),
				IsWishlisted: utils.Ptr(false),
			})
			if err != nil {
				return errors.WithMessage(err, "failed to upsert game detail record")
			}
		}
	}

	return nil
}
