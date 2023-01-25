package api

import (
	"fmt"
	db "vgtracker/backend/db"
	"vgtracker/backend/models"
)

type ProfileBackend struct{}

func (pb *ProfileBackend) UpdateProfile(input models.UpdateProfileInput) (*models.UpdateProfileOutput, error) {
	database, err := db.Connect()
	if err != nil {
		return nil, err
	}

	// return nil, fmt.Errorf("neat")

	_, err = database.Exec(`
		UPDATE profile 
		SET twitch_key = $1,
		psn_npsso = $2
	`, input.TwitchKey, input.PsnNpsso)
	if err != nil {
		return nil, fmt.Errorf("failed to update profile")
	}

	return &models.UpdateProfileOutput{}, nil
}
