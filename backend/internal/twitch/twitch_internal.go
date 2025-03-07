package twitch

import (
	"encoding/json"
	"fmt"
	"net/http"

	"vgtracker/backend/internal/config"

	"github.com/pkg/errors"
)

type TwitchInternalMethods interface {
	GetAccessToken(twitchSettings config.Twitch) (*AccessTokenResponse, error)
}

type TwitchInternalHandler struct {
	BaseUrl string
}

func NewTwitchInternal(BaseUrl string) TwitchInternalMethods {
	return &TwitchInternalHandler{
		BaseUrl,
	}
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (t *TwitchInternalHandler) GetAccessToken(twitchSettings config.Twitch) (*AccessTokenResponse, error) {

	resp, err := http.Post(fmt.Sprintf("%s?client_id=%s&client_secret=%s&grant_type=client_credentials", t.BaseUrl, twitchSettings.ClientID, twitchSettings.ClientSecret), "application/json", nil)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get twitch access token")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("failed to get twitch access token: %s", resp.Status)
	}

	var tokenResponse AccessTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, errors.WithMessage(err, "failed to decode twitch access token response")
	}

	return &tokenResponse, nil
}
