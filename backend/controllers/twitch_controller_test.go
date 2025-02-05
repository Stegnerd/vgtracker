package controllers_test

import (
	"errors"
	"vgtracker/backend/controllers"
	"vgtracker/backend/internal/config"
	"vgtracker/backend/internal/twitch"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TwitchController", func() {
	var twitchController controllers.TwitchControllerMethods

	var cfg *config.Config
	var twitchInternal *twitch.InternalMock

	BeforeEach(func() {
		cfg = &config.Config{
			Twitch: config.Twitch{
				ClientID:     "sample-clientID",
				ClientSecret: "sample-clientSecret",
			},
		}

		twitchInternal = &twitch.InternalMock{
			GetAccessTokenFunc: func(twitchConfig config.Twitch) (*twitch.AccessTokenResponse, error) {
				return &twitch.AccessTokenResponse{
					AccessToken: "twitchToken",
					ExpiresIn:   10,
					TokenType:   "grant",
				}, nil
			},
		}

		twitchController = controllers.NewTwitchController(cfg, twitchInternal)
	})

	Describe("GetTwitchAccessToken", func() {
		It("should return error when failing to get access token", func() {
			twitchInternal.GetAccessTokenFunc = func(twitchConfig config.Twitch) (*twitch.AccessTokenResponse, error) {
				return nil, errors.New("could not get access token")
			}

			token, err := twitchController.GetTwitchAccessToken()
			Expect(token).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(SatisfyAll(
				ContainSubstring("failed to get twitch access token"),
				ContainSubstring("could not get access token"),
			))
		})

		It("should succeed and return access token", func() {
			token, err := twitchController.GetTwitchAccessToken()
			Expect(token).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
