package twitch_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"vgtracker/backend/internal/config"
	"vgtracker/backend/internal/twitch"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TwitchInternal", func() {

	var twitchInternal twitch.TwitchInternalMethods

	var twitchSettings *config.Twitch
	var testServer *httptest.Server
	var validResponse *twitch.AccessTokenResponse
	baseUrl := "baseUrl"

	BeforeEach(func() {
		twitchSettings = &config.Twitch{
			ClientID:     "sample-clientID",
			ClientSecret: "sample-clientSecret",
		}

		validResponse = &twitch.AccessTokenResponse{
			AccessToken: "sampleAccessToken",
			ExpiresIn:   100,
			TokenType:   "accessGrant",
		}

		twitchInternal = twitch.NewTwitchInternal(baseUrl)
	})

	Context("GetAccessToken", func() {
		It("returns error when post fails", func() {

			testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				_, err := w.Write([]byte(`{ "response": "something went wrong" }`))
				Expect(err).To(Succeed())
			}))
			twitchInternal = twitch.NewTwitchInternal(testServer.URL)

			accessToken, err := twitchInternal.GetAccessToken(*twitchSettings)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(SatisfyAll(
				ContainSubstring("Internal Server Error"),
				ContainSubstring("failed to get twitch access token"),
			))
			Expect(accessToken).To(BeNil())
			defer testServer.Close()
		})

		It("returns error when decode fails", func() {
			testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write([]byte(`{ "response" b`))
				Expect(err).To(Succeed())
			}))
			twitchInternal = twitch.NewTwitchInternal(testServer.URL)

			accessToken, err := twitchInternal.GetAccessToken(*twitchSettings)
			Expect(accessToken).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(SatisfyAll(
				ContainSubstring("failed to decode twitch access token response"),
			))
			defer testServer.Close()
		})

		It("returns access token", func() {
			testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, err := fmt.Fprintln(w, `{"access_token":"sampleAccessToken","expires_in":100,"token_type":"accessGrant"}`)
				Expect(err).To(Succeed())
			}))
			twitchInternal = twitch.NewTwitchInternal(testServer.URL)

			accessToken, err := twitchInternal.GetAccessToken(*twitchSettings)
			Expect(err).ToNot(HaveOccurred())
			Expect(accessToken).ToNot(BeNil())
			Expect(accessToken).To(Equal(validResponse))
			defer testServer.Close()
		})
	})
})
