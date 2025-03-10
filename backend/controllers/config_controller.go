package controllers

import (
	"vgtracker/backend/internal/config"
	"vgtracker/backend/internal/utils"
	"vgtracker/backend/models"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
)

type ConfigControllerMethods interface {
	GetConfig() (*config.Config, error)
	UpdateConfig(input UpdateConfigInput) (*config.Config, error)
	RefreshConfig() error
}

type ConfigController struct {
	configInternalHandler config.ConfigInternalMethods
}

func NewConfigController(
	configInternalHandler config.ConfigInternalMethods,
) ConfigControllerMethods {

	return &ConfigController{
		configInternalHandler: configInternalHandler,
	}
}

func (c *ConfigController) GetConfig() (*config.Config, error) {
	appConfig, err := c.configInternalHandler.LoadConfig()
	if err != nil {
		return nil, errors.WithMessage(err, "could not load config")
	}

	return appConfig, nil
}

type UpdateConfigInput struct {
	Twitch       *config.Twitch       `json:"twitch"`
	Preset       *models.PresetConfig `json:"preset"`
	PrimaryColor *models.PaletteColor `json:"primaryColor"`
	SurfaceColor *models.SufaceColor  `json:"surfaceColor"`
	IsDarkTheme  *bool                `jons:"isDarkTheme"`
}

func (c *ConfigController) UpdateConfig(input UpdateConfigInput) (*config.Config, error) {
	appConfig, err := c.GetConfig()
	if err != nil {
		return nil, errors.WithMessage(err, "could not get config")
	}

	appConfig.Twitch = *utils.Patch(input.Twitch, appConfig.Twitch)
	appConfig.Theme.Preset = *utils.Patch(input.Preset, appConfig.Theme.Preset)
	appConfig.Theme.PrimaryColor = *utils.Patch(input.PrimaryColor, appConfig.Theme.PrimaryColor)
	appConfig.Theme.SurfaceColor = *utils.Patch(input.SurfaceColor, appConfig.Theme.SurfaceColor)
	appConfig.Theme.IsDarkTheme = *utils.Patch(input.IsDarkTheme, appConfig.Theme.IsDarkTheme)

	byteArr, err := toml.Marshal(appConfig)
	if err != nil {
		return nil, errors.WithMessage(err, "could not marshal config")
	}

	err = c.configInternalHandler.SaveConfig(byteArr)
	if err != nil {
		return nil, errors.WithMessage(err, "could not save config")
	}

	return appConfig, nil
}

func (c *ConfigController) RefreshConfig() error {
	return nil
}
