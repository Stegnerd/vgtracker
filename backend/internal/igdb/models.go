package igdb

type GameType string
type GenreType string
type PlatformType string

type VGTSearchResults struct {
	Items []VGTGame `json:"items"`
}

type VGTGame struct {
	GameID           int            `json:"id"`
	Name             string         `json:"name"`
	GameType         GameType       `json:"gameType"`
	Genres           []GenreType    `json:"genres"`
	Platforms        []PlatformType `json:"platforms"`
	ThumbnailURL     string         `json:"thumbnailURL"`
	CoverURL         string         `json:"coverURL"`
	FirstReleaseYear int            `json:"firstReleaseYear"`
	Summary          string         `json:"summary"`
	// values populated from user data (owned/played)
}

/*
TODO: potentiall in another package db models
*/
type VGTOwnderDetails struct {
	Owned  bool
	Played bool
	Beaten bool
	//MediaType
}
