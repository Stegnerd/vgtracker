package controllers

import (
	"vgtracker/backend/internal/steam"
)

type SteamControllerMethods interface {
}

type SteamController struct {
	steamInternalHandler steam.SteamInternalMethods
}

func NewSteamController(steamInternalHandler steam.SteamInternalMethods) SteamControllerMethods {
	return &SteamController{
		steamInternalHandler: steamInternalHandler,
	}
}
