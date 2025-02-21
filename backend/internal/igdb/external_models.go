package igdb

// SearchResultGame - model coming from IGDB
type IGDBSearchResultGame struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	GameType int    `json:"game_type"`
	Genres   []int  `json:"genres"`
}

func (IGDBSearchResultGame) convertGameType(incomingGameType int) GameType {
	switch incomingGameType {
	case 0:
		return MainGame
	case 1:
		return DLCAddon
	case 2:
		return Expansion
	case 3:
		return Bundle
	case 4:
		return StandaloneExpansion
	case 5:
		return Mod
	case 6:
		return Episode
	case 7:
		return Season
	case 8:
		return Remake
	case 9:
		return Remaster
	default:
		return GameTypeUnknown
	}
}

func (IGDBSearchResultGame) convertGenre(incomingResults IGDBSearchResultGame) []Genre {
	results := make([]Genre, len(incomingResults.Genres))

	for index, id := range incomingResults.Genres {
		switch id {
		case 2:
			results[index] = PointAndClick
		case 4:
			results[index] = Fighting
		case 5:
			results[index] = Shooter
		case 7:
			results[index] = Music
		case 8:
			results[index] = Platform
		case 9:
			results[index] = Puzzle
		case 10:
			results[index] = Racing
		case 11:
			results[index] = RTS
		case 12:
			results[index] = RPG
			//results[index] = RPG
		case 13:
			results[index] = Simulator
		case 14:
			results[index] = Sport
		case 15:
			results[index] = Strategy
		case 16:
			results[index] = TurnBasedStrategy
		case 24:
			results[index] = Tactical
		case 25:
			results[index] = HackNSlashBeatMUp
		case 26:
			results[index] = QuizTrivia
		case 30:
			results[index] = Pinball
		case 31:
			results[index] = Adventure
		case 32:
			results[index] = Indie
		case 33:
			results[index] = Arcade
		case 34:
			results[index] = VisualNovel
		case 35:
			results[index] = CardOrBordGame
		case 36:
			results[index] = Moba
		default:
			results[index] = GenreUnknown
		}
	}

	return results
}
