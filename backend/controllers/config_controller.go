package controllers

import (
	"vgtracker/backend/internal/config"

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
	Twitch config.Twitch `json:"twitch"`
}

func (c *ConfigController) UpdateConfig(input UpdateConfigInput) (*config.Config, error) {
	appConfig, err := c.GetConfig()
	if err != nil {
		return nil, errors.WithMessage(err, "could not get config")
	}

	appConfig.Twitch = input.Twitch

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
