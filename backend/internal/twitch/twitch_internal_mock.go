package twitch

import "vgtracker/backend/internal/config"

type InternalMock struct {
	GetAccessTokenFunc func(twitchConfig config.Twitch) (*AccessTokenResponse, error)
}

var _ TwitchInternalMethods = &InternalMock{}

func (i InternalMock) GetAccessToken(twitchSettings config.Twitch) (*AccessTokenResponse, error) {
	return i.GetAccessTokenFunc(twitchSettings)
}
