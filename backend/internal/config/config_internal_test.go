package config_test

import (
	"vgtracker/backend/internal/config"
	"vgtracker/backend/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/afero"
)

var _ = Describe("ConfigInternal", func() {

	var configInternal config.ConfigInternalMethods

	var memoryFS afero.Fs
	var sampleConfig *config.Config

	BeforeEach(func() {
		memoryFS = afero.NewMemMapFs()

		configInternal = config.NewConfigInternal(memoryFS)
	})

	initExistingConfig := func() {
		sampleConfig = &config.Config{
			Twitch: config.Twitch{
				ClientID:     "sample-clientID",
				ClientSecret: "sample-clientSecret",
			},
		}
		configStr, err := toml.Marshal(sampleConfig)
		Expect(err).ToNot(HaveOccurred())

		err = memoryFS.MkdirAll("config.toml", 0755)
		Expect(err).ToNot(HaveOccurred())
		err = afero.WriteFile(memoryFS, "config.toml", configStr, 0644)
		Expect(err).ToNot(HaveOccurred())
	}

	Context("NewConfig", func() {
		When("config directory exists", func() {
			It("creates new config and returns", func() {
				initExistingConfig()
				conf, err := configInternal.NewConfig()
				Expect(err).ToNot(HaveOccurred())
				Expect(conf).To(Equal(sampleConfig))
			})
		})

		When("config directory does not exist", func() {
			It("creates new config and returns", func() {
				conf, err := configInternal.NewConfig()
				Expect(err).ToNot(HaveOccurred())
				Expect(conf).To(Equal(&config.Config{
					Twitch: config.Twitch{},
					Theme: config.ThemeSettings{
						Preset:       models.Lara,
						PrimaryColor: models.Emerald,
						SurfaceColor: models.Zinc,
						IsDarkTheme:  true,
					},
				}))
			})
		})
	})

	Context("LoadConfig", func() {
		BeforeEach(func() {
			initExistingConfig()
		})

		It("returns config if successful", func() {
			conf, err := configInternal.LoadConfig()
			Expect(err).ToNot(HaveOccurred())
			Expect(conf).To(Equal(sampleConfig))
		})
	})

	Context("SaveConfig", func() {
		BeforeEach(func() {
			initExistingConfig()
		})

		It("updates config file and returns nil", func() {
			sampleConfig.Twitch.ClientID = "id-updated"
			sampleConfig.Twitch.ClientSecret = "secret-updated"

			updatedConfigStr, err := toml.Marshal(sampleConfig)
			Expect(err).ToNot(HaveOccurred())

			err = configInternal.SaveConfig(updatedConfigStr)
			Expect(err).ToNot(HaveOccurred())
			Expect(err).To(BeNil())
		})
	})
})
