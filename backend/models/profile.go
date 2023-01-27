package models

type Profile struct {
	TwitchKey string `db:"twitch_key" json:"twitchKey"`
	PsnNpsso  string `db:"psn_npsso" json:"psnNpsso"`
}

// API MODELS

type ReadProfileOutput struct {
	UserProfile Profile `json:"userProfile"`
}

type UpdateProfileInput struct {
	TwitchKey string `json:"twitchKey"`
	PsnNpsso  string `json:"psnNpsso"`
}

type UpdateProfileOutput struct{}
