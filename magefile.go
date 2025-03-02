//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	// mg contains helpful utility functions, like Deps
	_ "modernc.org/sqlite"
)

const (
	migrationsDir = "./backend/internal/db/migrations"
)

type DB mg.Namespace

// MigrateUp runs all pending database migrations.
func (DB) MigrateUp() error {
	dbFile, err := getDBUrl()
	if err != nil {
		return err
	}

	fmt.Println("Running migrations up...")
	return sh.RunV("migrate", "-path", migrationsDir, "-database", dbFile, "up")
}

// MigrateDown reverts the most recent migration.
func (DB) MigrateDown() error {
	dbFile, err := getDBUrl()
	if err != nil {
		return err
	}

	return sh.RunV("migrate", "-path", migrationsDir, "-database", dbFile, "down", "1")
}

// MigrateCreate creates a new migration file.
func (DB) MigrateCreate(name string) error {
	fmt.Printf("Creating migration: %s\n", name)
	return sh.RunV("migrate", "create", "-ext", "sql", "-dir", migrationsDir, "-seq", name)
}

func getDBUrl() (string, error) {
	appConfigDir, err := os.UserHomeDir()
	if err != nil {
		panic("could not get user config directory")
	}

	if _, err := os.Stat(appConfigDir + "/vgtracker"); os.IsNotExist(err) {
		dirErr := os.Mkdir(appConfigDir+"/vgtracker", 0777)
		if dirErr != nil {
			panic("could not create config directory")
		}
	}

	url := fmt.Sprintf("%s://%s/vgtracker/vgtracker.db", "sqlite3", appConfigDir)

	return url, nil
}
