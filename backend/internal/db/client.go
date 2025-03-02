package db

import (
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"path/filepath"
	"vgtracker/backend/internal/utils"

	"github.com/pkg/errors"
	"github.com/spf13/afero"

	// unnamed import to reguest sql driver
	_ "modernc.org/sqlite"
)

const DbFileName = "vgtracker.db"

//go:embed migrations/*.sql
var schemaFs embed.FS

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
			return nil, err
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

	if err = runMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DB) makeFileWithDirectory() error {

	_, err := d.AppFS.Create(DbFileName)
	if err != nil {
		return errors.WithMessage(err, "could not create app config file")
	}

	return nil
}

func runMigrations(db *sql.DB) error {
	source, err := iofs.New(schemaFs, "migrations")
	if err != nil {
		panic(errors.WithMessage(err, "could not create database driver"))
	}

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		panic(errors.WithMessage(err, "could not create database instance"))
	}

	m, err := migrate.NewWithInstance("iofs", source, "sqlite3", instance)
	if err != nil {
		panic(errors.WithMessage(err, "could not create database migration"))
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(errors.WithMessage(err, "could not run migration"))
	}

	return nil
}
