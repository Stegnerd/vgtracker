package controllers

import "vgtracker/backend/internal/gamedetails"

type GameDetailMethods interface {
	GetGameDetails(igdbID int) (*gamedetails.GetGameDetailOutput, error)
	UpsertGameDetails(input gamedetails.UpsertGameDetailInput) (*gamedetails.GetGameDetailOutput, error)
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

// UpsertGameDetails implements GameDetailMethods.
func (g *GameDetailController) UpsertGameDetails(input gamedetails.UpsertGameDetailInput) (*gamedetails.GetGameDetailOutput, error) {
	return g.gameDetailInternalHandler.UpsertGameDetailRecord(input)
}
