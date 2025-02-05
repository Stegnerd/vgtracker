package backend

import (
	"context"
	"fmt"
	"vgtracker/backend/controllers"
	"vgtracker/backend/internal/config"
	"vgtracker/backend/internal/twitch"
	"vgtracker/backend/internal/utils"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

const (
	TwitchAuthURL = "https://id.twitch.tv/oauth2/token"
)

// App struct
type App struct {
	ctx context.Context
	// configs
	AppFS afero.Fs
	Cfg   *config.Config
	// controllers
	ConfigController controllers.ConfigControllerMethods
	TwitchController controllers.TwitchControllerMethods
}

// NewApp creates a new App application struct
func NewApp() *App {
	configDir, err := utils.GetConfigDirectory()
	if err != nil {
		panic(errors.WithMessage(err, "could not get config directory"))
	}

	appFS := afero.NewBasePathFs(afero.NewOsFs(), configDir)

	// config setup
	configInternal := config.NewConfigInternal(appFS)
	configController := controllers.NewConfigController(configInternal)
	appConfig, err := configInternal.NewConfig()
	if err != nil {
		panic(errors.WithMessage(err, "could not create or load config"))
	}

	// twitch setup
	twitchInternal := twitch.NewTwitchInternal(TwitchAuthURL)
	twitchController := controllers.NewTwitchController(appConfig, twitchInternal)

	var twitchToken *twitch.AccessTokenResponse
	if appConfig.Twitch.ClientID != "" && appConfig.Twitch.ClientSecret != "" {
		twitchToken, err = twitchInternal.GetAccessToken(appConfig.Twitch)
		if err != nil {
			panic(errors.WithMessage(err, "could not get twitch token"))
		}
	}
	fmt.Printf("\n twitchToken: %+v  \n", twitchToken)

	return &App{
		Cfg: appConfig,
		//

		//
		ConfigController: configController,
		TwitchController: twitchController,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Start(ctx context.Context) {
	a.startup(ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
