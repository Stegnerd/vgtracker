package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vgtracker/backend/internal/config"

	"github.com/pkg/errors"
)

type SteamInternalMethods interface {
	GetOwnedGames(steamID string) (*OwnedGamesResponse, error)
}

type SteamInternalHandler struct {
	config *config.Config
}

func NewSteamInternal(config *config.Config) SteamInternalMethods {
	return &SteamInternalHandler{
		config: config,
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

func (s *SteamInternalHandler) GetOwnedGames(steamID string) (*OwnedGamesResponse, error) {
	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json&include_appinfo=true",
		s.config.Steam.APIKey,
		steamID,
	)

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
