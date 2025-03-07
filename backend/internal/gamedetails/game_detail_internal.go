package gamedetails

import (
	"database/sql"
	"errors"
	"time"
)

type GameDetailInternalMethods interface {
	GetGameDetailRecord(id int) (*GetGameDetailOutput, error)
	UpsertGameDetailRecord()
}

type GameDetailInternalHandler struct {
	DB *sql.DB
}

func NewGameDetailInternal(
	DB *sql.DB,
) GameDetailInternalMethods {
	return &GameDetailInternalHandler{
		DB,
	}
}

type GetGameDetailOutput struct {
	ID           int       `json:"id" db:"id"`
	IGDBID       int       `json:"igdbID" db:"igdb_id"`
	IsOwned      bool      `json:"isOwned" db:"is_owned"`
	IsBeaten     bool      `json:"isBeaten" db:"is_beaten"`
	IsWishlisted bool      `json:"isWishlisted" db:"is_wishlisted"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

// GetBacklogRecord implements BacklogInternalMethods.
func (gd *GameDetailInternalHandler) GetGameDetailRecord(id int) (*GetGameDetailOutput, error) {
	if id == 0 {
		return nil, errors.New("id must be provided to get details")
	}

	query := `
		SELECT id,
					 igdb_id,
					 is_owned,
					 is_beaten,
					 is_wishlisted,
					 created_at,
					 updated_at
		FROM game_details
		WHERE igdb_id = ?
	`
	output := &GetGameDetailOutput{}
	err := gd.DB.QueryRow(query, id).Scan(
		&output.ID,
		&output.IGDBID,
		&output.IsOwned,
		&output.IsBeaten,
		&output.IsWishlisted,
		&output.CreatedAt,
		&output.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return output, nil
		}
		return nil, err
	}

	return output, nil
}

type UpserGameDetailInput struct {
	ID string
}

// UpsertGameDetailRecord implements GameDetailInternalMethods.
func (gd *GameDetailInternalHandler) UpsertGameDetailRecord() {
	panic("unimplemented")
}
