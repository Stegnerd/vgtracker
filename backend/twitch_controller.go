package backend

import (
	"vgtracker/backend/internal/config"
	"vgtracker/backend/internal/twitch"

	"github.com/pkg/errors"
)

type TwitchControllerMethods interface {
	GetTwitchAccessToken() (*twitch.AccessTokenResponse, error)
}

type TwitchController struct {
	config                *config.Config
	twitchInternalHandler twitch.TwitchInternalMethods
}

func NewTwitchController(config *config.Config, twitchInternalHandler twitch.TwitchInternalMethods) TwitchControllerMethods {
	return &TwitchController{
		config:                config,
		twitchInternalHandler: twitchInternalHandler,
	}
}

func (t *TwitchController) GetTwitchAccessToken() (*twitch.AccessTokenResponse, error) {
	token, err := t.twitchInternalHandler.GetAccessToken(t.config.Twitch)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get twitch access token")
	}

	return token, nil
}
