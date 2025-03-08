package controllers

import "vgtracker/backend/internal/gamedetails"

type GameDetailMethods interface {
	GetGameDetails(id int) (*gamedetails.GetGameDetailOutput, error)
}

type GameDetailController struct {
	gameDetailInternalHandler gamedetails.GameDetailInternalMethods
}

func NewGameDetailController(
	gameDetailInternalHandler gamedetails.GameDetailInternalMethods,
) GameDetailMethods {
	return &GameDetailController{
		gameDetailInternalHandler,
	}
}

// GetGameDetails implements GameDetailMethods.
func (g *GameDetailController) GetGameDetails(id int) (*gamedetails.GetGameDetailOutput, error) {
	return g.gameDetailInternalHandler.GetGameDetailRecord(id)
}
