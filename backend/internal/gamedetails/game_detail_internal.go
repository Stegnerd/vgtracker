package gamedetails

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"vgtracker/backend/internal/utils"
)

type GameDetailInternalMethods interface {
	GetGameDetailRecord(id int) (*GetGameDetailOutput, error)
	UpsertGameDetailRecord(input UpsertGameDetailInput) (*GetGameDetailOutput, error)
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

type UpsertGameDetailInput struct {
	ID           *int
	IGDBID       *int
	IsOwned      *bool `json:"isOwned"`
	IsBeaten     *bool `json:"isBeaten"`
	IsWishlisted *bool `json:"isWishlisted"`
}

// UpsertGameDetailRecord implements GameDetailInternalMethods.
func (gd *GameDetailInternalHandler) UpsertGameDetailRecord(input UpsertGameDetailInput) (*GetGameDetailOutput, error) {

	if input.IGDBID == nil {
		return nil, errors.New("must provide IGDBID in order to save details.")
	}

	if input.ID == nil {
		return gd.create(input)
	}

	return gd.update(input)
}

func (gd *GameDetailInternalHandler) create(input UpsertGameDetailInput) (*GetGameDetailOutput, error) {
	var query string
	var args []interface{}
	args = append(args, input.IGDBID)

	query = `INSERT INTO game_detail(igdb_id, created_at, updated_at, is_owned, is_beaten, is_wishlisted)
					 VALUES(?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?, ?, ?)`

	// TODO: BRING PATCHING UP TO THE CONTROLLER LEVELS
	input.IsOwned = utils.Patch(input.IsOwned, false)
	input.IsBeaten = utils.Patch(input.IsBeaten, false)
	input.IsWishlisted = utils.Patch(input.IsWishlisted, false)

	args = append(args, *input.IsOwned, *input.IsBeaten, *input.IsWishlisted)

	insertResult, err := gd.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to insert record for game: '%d' due to %w", *input.IGDBID, err)
	}

	newID, err := insertResult.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get new id after insert: %w", err)
	}

	return gd.GetGameDetailRecord(int(newID))
}

func (gd *GameDetailInternalHandler) update(input UpsertGameDetailInput) (*GetGameDetailOutput, error) {
	var query string
	var args []interface{}

	currentModel, err := gd.GetGameDetailRecord(*input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to read current record for id: '%d' to update", *input.ID)
	}

	query = `UPDATE game_detail
					 SET updated_at = CURRENT_TIMESTAMP,
					   is_owned = ?,
						 is_beaten = ?,
						 is_wishlisted = ?
					 WHERE id = ? and igdb_id = ?
					 `

	input.IsOwned = utils.Patch(input.IsOwned, currentModel.IsOwned)
	input.IsBeaten = utils.Patch(input.IsBeaten, currentModel.IsBeaten)
	input.IsWishlisted = utils.Patch(input.IsWishlisted, currentModel.IsWishlisted)

	args = append(args, *input.IsOwned, *input.IsBeaten, *input.IsWishlisted, *input.ID, *input.IGDBID)

	insertResult, err := gd.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to insert record for game: '%d' due to %w", *input.IGDBID, err)
	}

	newID, err := insertResult.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get new id after insert: %w", err)
	}

	return gd.GetGameDetailRecord(int(newID))
}
