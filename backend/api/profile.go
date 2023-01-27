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

func (pb *ProfileBackend) ReadProfile() (*models.ReadProfileOutput, error) {
	database, err := db.Connect()
	if err != nil {
		return nil, err
	}

	var profile models.Profile
	err = database.Get(&profile, `
		SELECT twitch_key,
		       psn_npsso
		FROM profile`)
	if err != nil {
		return nil, fmt.Errorf("failed to read profile")
	}

	return &models.ReadProfileOutput{
		UserProfile: profile,
	}, nil
}
