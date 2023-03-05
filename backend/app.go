package backend

import (
	"context"
	"database/sql"
	"fmt"
	"vgtracker/backend/api"
	"vgtracker/backend/db"
	"vgtracker/backend/igdbwrapper"
)

// App struct
type App struct {
	ctx        context.Context
	db         *sql.DB
	igdbClient *igdbwrapper.IGDBWrapper
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup - is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.db = db.InitDB()

	igdbClient := igdbwrapper.NewIGDBWrapperService()
	profileService := api.NewProfileBackendService()
	profile, err := profileService.ReadProfile()
	if err != nil {
		panic("Failed to start to up db was not in a good state")
	}
	fmt.Printf("\n\n starting the get call \n\n")
	if profile.UserProfile.TwitchKey != "" && profile.UserProfile.TwitchSecret != "" {
		// get access token
		igdbClient.GetTwitchAccessToken(profile.UserProfile.TwitchKey, profile.UserProfile.TwitchSecret)
		igdbClient.NewClient()
	}
	igdbClient.Test()

	a.igdbClient = igdbClient
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
