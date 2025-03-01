//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"

	// mg contains helpful utility functions, like Deps
	_ "modernc.org/sqlite"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// // A build step that requires additional params, or platform specific steps for example
// func Build() error {
// 	mg.Deps(InstallDeps)
// 	fmt.Println("Building...")
// 	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
// 	return cmd.Run()
// }

// // A custom install step if you need your bin someplace other than go/bin
// func Install() error {
// 	mg.Deps(Build)
// 	fmt.Println("Installing...")
// 	return os.Rename("./MyApp", "/usr/bin/MyApp")
// }

// // Manage your deps, or running package managers.
// func InstallDeps() error {
// 	fmt.Println("Installing Deps...")
// 	cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
// 	return cmd.Run()
// }

// // Clean up after yourself
// func Clean() {
// 	fmt.Println("Cleaning...")
// 	os.RemoveAll("MyApp")
// }

const (
	migrationsDir = "./backend/internal/db/migrations"
)

// MigrateUp runs all pending database migrations.
func MigrateUp() error {
	dbFile, err := getDBUrl()
	if err != nil {
		return err
	}

	fmt.Println("Running migrations up...")
	return sh.RunV("migrate", "-path", migrationsDir, "-database", dbFile, "up")
}

// MigrateDown reverts the most recent migration.
func MigrateDown() error {
	dbFile, err := getDBUrl()
	if err != nil {
		return err
	}

	fmt.Println("Running migrations down...")
	return sh.RunV("migrate", "-path", migrationsDir, "-database", dbFile, "down")
}

// MigrateCreate creates a new migration file.
func MigrateCreate(name string) error {
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
