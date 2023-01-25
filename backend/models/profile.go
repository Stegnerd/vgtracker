package models

type Profile struct {
	TwitchKey string `db:"twitch_key" json:""`
}

// API MODELS
type UpdateProfileInput struct {
	TwitchKey string `json:"twitchKey"`
	PsnNpsso  string `json:"psnNpsso"`
}

type UpdateProfileOutput struct{}
