package igdb

type GameType string
type Genre string

const (
	MainGame            GameType = "Main Game"
	DLCAddon            GameType = "DLC Addon"
	Expansion           GameType = "Expansion"
	Bundle              GameType = "Bundle"
	StandaloneExpansion GameType = "Standalone Expansion"
	Mod                 GameType = "Mod"
	Episode             GameType = "Episode"
	Season              GameType = "Season"
	Remake              GameType = "Remake"
	Remaster            GameType = "Remaster"
	GameTypeUnknown     GameType = "Unknown"
)

const (
	PointAndClick     Genre = "Point And Click"
	Fighting          Genre = "Fighting"
	Shooter           Genre = "Shooter"
	Music             Genre = "Music"
	Platform          Genre = "Platformer"
	Puzzle            Genre = "Puzzle"
	Racing            Genre = "Racing"
	RTS               Genre = "Real Time Strategy (RTS)"
	RPG               Genre = "Role Playing (RPG)"
	Simulator         Genre = "Simulator"
	Sport             Genre = "Sport"
	Strategy          Genre = "Strategy"
	TurnBasedStrategy Genre = "Turn Based Strategy (TBS)"
	Tactical          Genre = "Tactical"
	HackNSlashBeatMUp Genre = "Hack N' Slash/ Beat'em Up"
	QuizTrivia        Genre = "Quiz/Trivia"
	Pinball           Genre = "Pinball"
	Adventure         Genre = "Adventure"
	Indie             Genre = "Indie"
	Arcade            Genre = "Arcade"
	VisualNovel       Genre = "Visual Novel"
	CardOrBordGame    Genre = "Card Or Board Game"
	Moba              Genre = "MOBA"
	GenreUnknown      Genre = "Unknown"
)

type VGTSearchResults struct {
	Items []VGTGame `json:"items"`
}

type VGTGame struct {
	Name     string   `json:"name"`
	GameType GameType `json:"gameType"`
	Genres   []Genre  `json:"genres"`
}
