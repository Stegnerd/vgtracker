package config

import (
	"vgtracker/backend/internal/utils"
	"vgtracker/backend/models"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

const configFileName = "config.toml"

type Config struct {
	Twitch Twitch        `json:"twitch"`
	Steam  Steam         `json:"steam"`
	Theme  ThemeSettings `json:"theme"`
}

type Twitch struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

type Steam struct {
	SteamID string `json:"steamID"`
	APIKey  string `json:"apiKey"`
}

type ThemeSettings struct {
	Preset       models.PresetConfig `json:"preset"`
	PrimaryColor models.PaletteColor `json:"primaryColor"`
	SurfaceColor models.SufaceColor  `json:"surfaceColor"`
	IsDarkTheme  bool                `jons:"isDarkTheme"`
}

type ConfigInternalMethods interface {
	NewConfig() (*Config, error)
	LoadConfig() (*Config, error)
	SaveConfig(saveArr []byte) error
}

type ConfigInternal struct {
	AppFS afero.Fs
}

func NewConfigInternal(AppFS afero.Fs) ConfigInternalMethods {
	return &ConfigInternal{
		AppFS: AppFS,
	}
}

func (c *ConfigInternal) NewConfig() (*Config, error) {
	if utils.FileExists(c.AppFS, configFileName) {
		return c.LoadConfig()
	} else {
		return c.makeConfigWithDirectory()
	}
}

func (c *ConfigInternal) makeConfigWithDirectory() (*Config, error) {

	// TODO: not sure if discarding file auto closes
	_, err := c.AppFS.Create(configFileName)
	if err != nil {
		return nil, errors.WithMessage(err, "could not create app config file")
	}

	defaultConfig := newDefaultConfig()

	byteArr, err := toml.Marshal(defaultConfig)
	if err != nil {
		return nil, errors.WithMessage(err, "could not marshal default config")
	}

	err = afero.WriteFile(c.AppFS, configFileName, byteArr, 0777)
	if err != nil {
		return nil, errors.WithMessage(err, "could not write to config file")
	}

	return defaultConfig, nil
}

func (c *ConfigInternal) LoadConfig() (*Config, error) {
	configFile, err := c.AppFS.Open(configFileName)
	if err != nil {
		return nil, errors.WithMessage(err, "could not open config file")
	}
	defer configFile.Close()

	bytes, err := afero.ReadAll(configFile)
	if err != nil {
		return nil, errors.WithMessage(err, "could not read config file")
	}

	var config *Config
	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, errors.WithMessage(err, "could not unmarshal config file")
	}

	return config, nil
}

func (c *ConfigInternal) SaveConfig(saveArr []byte) error {
	err := afero.WriteFile(c.AppFS, configFileName, saveArr, 0777)
	if err != nil {
		return errors.WithMessage(err, "could not write to config file")
	}

	return nil
}

func newDefaultConfig() *Config {
	return &Config{
		Twitch: Twitch{
			ClientID:     "",
			ClientSecret: "",
		},
		Steam: Steam{
			SteamID: "",
			APIKey:  "",
		},
		Theme: ThemeSettings{
			Preset:       models.Lara,
			PrimaryColor: models.Emerald,
			SurfaceColor: models.Zinc,
			IsDarkTheme:  true,
		},
	}
}
