package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

//go:embed migrations
var fs embed.FS

var db *sql.DB

const DB_NAME = "vgtracker.db"

func InitDB() *sql.DB {
	ensureDBExists()

	db, err := sql.Open("sqlite", DB_NAME)
	if err != nil {
		fmt.Println("db open err")
		log.Fatalln(err)
	}
	if err := ensureSchema(db); err != nil {
		fmt.Println("migration failed")
		log.Fatal(err)
	}

	return db
}

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite", DB_NAME)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to db")
	}

	return db, nil
}

func ensureDBExists() {

	_, err := os.Stat("vgtracker.db")
	if os.IsNotExist(err) {
		file, err := os.Create("vgtracker.db")
		if err != nil {
			log.Println("Failed to create db file")
			log.Fatal(err)
		}
		file.Close()
	}
}

func ensureSchema(db *sql.DB) error {
	sourceInstance, err := iofs.New(fs, "migrations")
	if err != nil {
		return fmt.Errorf("invalid source instance, %w", err)
	}
	targetInstance, err := sqlite.WithInstance(db, new(sqlite.Config))
	if err != nil {
		return fmt.Errorf("invalid target sqlite instance, %w", err)
	}
	m, err := migrate.NewWithInstance("iofs", sourceInstance, "sqlite", targetInstance)
	if err != nil {
		return fmt.Errorf("failed to initialize migrate instance, %w", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}
