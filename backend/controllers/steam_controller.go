package controllers

import (
	"vgtracker/backend/internal/steam"
)

type SteamControllerMethods interface {
	SyncOwnedGames() error
}

type SteamController struct {
	steamInternalHandler steam.SteamInternalMethods
}

// SyncOwnedGames implements SteamControllerMethods.
func (s *SteamController) SyncOwnedGames() error {
	return s.steamInternalHandler.SyncOwnedGames()
}

func NewSteamController(steamInternalHandler steam.SteamInternalMethods) SteamControllerMethods {
	return &SteamController{
		steamInternalHandler: steamInternalHandler,
	}
}
