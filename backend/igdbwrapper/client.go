package igdbwrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/labstack/gommon/log"
	"net/http"
	"vgtracker/backend/models"
)

const TWITCH_URL = "https://id.twitch.tv/oauth2/token"

type IGDBWrapper struct {
	ClientID *string
	Client   *igdb.Client
	Token    *models.TwitchApiResult
}

type IGDBWrapperService interface {
	NewClient() *igdb.Client
}

func NewIGDBWrapperService() *IGDBWrapper {
	return &IGDBWrapper{}
}

func (w *IGDBWrapper) GetTwitchAccessToken(id, secret string) {
	if w.Token != nil && w.Token.AccessToken != "" {
		// check expiration
		// if not expired return it
		// return w.Token

		// else keep going
	}

	endpoint := fmt.Sprintf("%s?client_id=%s&client_secret=%s&grant_type=client_credentials", TWITCH_URL, id, secret)
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Errorf("Failed to get access token", err)
	}
	defer resp.Body.Close()

	var twitchResponse *models.TwitchApiResult
	if err = json.NewDecoder(resp.Body).Decode(&twitchResponse); err != nil {
		log.Fatalf("Failed to decode response body: %s", err.Error())
	}

	fmt.Printf("Here is the twitchResponse: %+v", twitchResponse)

	w.ClientID = &id
	w.Token = twitchResponse
}

func (w *IGDBWrapper) NewClient() *igdb.Client {
	client := igdb.NewClient(*w.ClientID, w.Token.AccessToken, nil)
	fmt.Printf("\n\n got the new client \n\n")

	w.Client = client
	return client
}

func (w *IGDBWrapper) Test() {
	//tester, err := w.Client.Games.Search("megaman")
	byPop := igdb.ComposeOptions(
		igdb.SetLimit(5),
		igdb.SetFields("name", "cover"),
		igdb.SetOrder("hypes", igdb.OrderDescending),
		igdb.SetFilter("category", igdb.OpEquals, "0"),
		igdb.SetFilter("cover", igdb.OpNotEquals, "null"),
	)

	// Retrieve PS4 inter-console exclusives
	PS4, err := w.Client.Games.Index(
		byPop, // top 5 popular results
	)
	if err != nil {
		fmt.Printf("Here is the error: %s", err.Error())
	}

	test := PS4[0]
	fmt.Printf("\n\n Here is the response for the call %+v \n\n", test)
}

func isTokenExpired() bool {
	return false
}
