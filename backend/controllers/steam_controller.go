package controllers

import (
	"vgtracker/backend/internal/steam"

	"github.com/pkg/errors"
)

type SteamControllerMethods interface {
	GetOwnedGames(steamID string) (*steam.OwnedGamesResponse, error)
}

type SteamController struct {
	steamInternalHandler steam.SteamInternalMethods
}

func NewSteamController(steamInternalHandler steam.SteamInternalMethods) SteamControllerMethods {
	return &SteamController{
		steamInternalHandler: steamInternalHandler,
	}
}

func (s *SteamController) GetOwnedGames(steamID string) (*steam.OwnedGamesResponse, error) {
	games, err := s.steamInternalHandler.GetOwnedGames(steamID)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get steam owned games")
	}

	return games, nil
}
