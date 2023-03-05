package api

import (
	"fmt"
	db "vgtracker/backend/db"
	"vgtracker/backend/models"
)

type ProfileBackendService interface {
	UpdateProfile(input models.UpdateProfileInput) (*models.UpdateProfileOutput, error)
	ReadProfile() (*models.ReadProfileOutput, error)
}

func NewProfileBackendService() ProfileBackendService {
	return &ProfileBackend{}
}

type ProfileBackend struct{}

func (pb *ProfileBackend) UpdateProfile(input models.UpdateProfileInput) (*models.UpdateProfileOutput, error) {
	database, err := db.Connect()
	if err != nil {
		return nil, err
	}

	_, err = database.Exec(`
		UPDATE profile 
		SET twitch_key = $1,
		    twitch_secret = $2,
			psn_npsso = $3
	`, input.TwitchKey, input.TwitchSecret, input.PsnNpsso)
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
		       twitch_secret,
		       psn_npsso
		FROM profile`)
	if err != nil {
		return nil, fmt.Errorf("failed to read profile")
	}

	return &models.ReadProfileOutput{
		UserProfile: profile,
	}, nil
}
