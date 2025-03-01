package db

import (
	"database/sql"
	"path/filepath"
	"vgtracker/backend/internal/utils"

	"github.com/pkg/errors"
	"github.com/spf13/afero"

	// unnamed import to reguest sql driver
	_ "modernc.org/sqlite"
)

const DbFileName = "vgtracker.db"

type DBMethods interface {
	NewDB() (*sql.DB, error)
}

type DB struct {
	AppFS afero.Fs
}

func NewDBClient(AppFS afero.Fs) DBMethods {
	return &DB{
		AppFS,
	}
}

func (d *DB) NewDB() (*sql.DB, error) {
	if !utils.FileExists(d.AppFS, DbFileName) {
		err := d.makeFileWithDirectory()
		if err != nil {
			return err
		}
	}

	dir, err := utils.GetConfigDirectory()
	if err != nil {
		return nil, err
	}

	dbFilePath := filepath.Join(dir, DbFileName)

	db, err := sql.Open("sqlite", dbFilePath)
	if err != nil {
		panic(errors.WithMessage(err, "could not create or load db"))
	}

}

func (d *DB) makeFileWithDirectory() error {

	_, err := d.AppFS.Create(DbFileName)
	if err != nil {
		return errors.WithMessage(err, "could not create app config file")
	}

	return nil
}
