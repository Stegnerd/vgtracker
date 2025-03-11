package controllers_test

import (
	"errors"
	"vgtracker/backend/controllers"
	"vgtracker/backend/internal/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ConfigController", func() {

	var configController controllers.ConfigControllerMethods

	var cfg *config.Config
	var configInternal *config.InternalMock

	BeforeEach(func() {

		cfg = &config.Config{
			Twitch: config.Twitch{
				ClientID:     "sample-clientID",
				ClientSecret: "sample-clientSecret",
			},
		}

		configInternal = &config.InternalMock{
			LoadConfigFunc: func() (*config.Config, error) {
				return cfg, nil
			},
			SaveConfigFunc: func(saveArr []byte) error {
				return nil
			},
		}

		configController = controllers.NewConfigController(configInternal)
	})

	Describe("GetConfig", func() {
		It("should return error when failing to load config", func() {
			configInternal.LoadConfigFunc = func() (*config.Config, error) {
				return nil, errors.New("load conf err")
			}

			conf, err := configController.GetConfig()
			Expect(conf).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(SatisfyAll(
				ContainSubstring("load conf err"),
				ContainSubstring("could not load config"),
			))
		})

		It("should succeed and return config", func() {
			conf, err := configController.GetConfig()
			Expect(conf).To(Equal(cfg))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("UpdateConfig", func() {
		It("should return error when failing to load config", func() {
			configInternal.LoadConfigFunc = func() (*config.Config, error) {
				return nil, errors.New("load conf err")
			}

			conf, err := configController.UpdateConfig(controllers.UpdateConfigInput{})
			Expect(conf).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(SatisfyAll(
				ContainSubstring("load conf err"),
				ContainSubstring("could not load config"),
			))
		})

		It("should return error when failing to save config", func() {
			configInternal.SaveConfigFunc = func([]byte) error {
				return errors.New("save conf err")
			}

			conf, err := configController.UpdateConfig(controllers.UpdateConfigInput{})
			Expect(conf).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(SatisfyAll(
				ContainSubstring("save conf err"),
				ContainSubstring("could not save config"),
			))
		})

		It("should succeed and return updated config", func() {
			conf, err := configController.UpdateConfig(controllers.UpdateConfigInput{
				Twitch: &config.Twitch{
					ClientID:     "updated-clientID",
					ClientSecret: "updated-clientSecret",
				},
			})
			Expect(err).To(BeNil())
			Expect(conf).To(Equal(&config.Config{
				Twitch: config.Twitch{
					ClientID:     "updated-clientID",
					ClientSecret: "updated-clientSecret",
				},
			}))
		})
	})
})
