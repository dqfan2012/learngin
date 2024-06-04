package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dqfan2012/learngin/pkg/db"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	migrateUp := flag.Bool("up", false, "Apply all up migrations")
	migrateDown := flag.Bool("down", false, "Apply all down migrations")
	makeMigration := flag.String("new", "", "Create new migration files with the given name")
	flag.Parse()

	if *makeMigration != "" {
		createMigrationFiles(*makeMigration)
		return
	}

	// Initialize the database
	db.Init()

	if *migrateUp {
		runMigrations("up")
	}

	if *migrateDown {
		runMigrations("down")
	}
}

func createMigrationFiles(name string) {
	timestamp := time.Now().Format("20060102150405")
	upFileName := fmt.Sprintf("migrations/%s_%s.up.sql", timestamp, name)
	downFileName := fmt.Sprintf("migrations/%s_%s.down.sql", timestamp, name)

	upFile, err := os.Create(upFileName)
	if err != nil {
		log.Fatalf("Failed to create up migration file: %v", err)
	}
	defer upFile.Close()

	downFile, err := os.Create(downFileName)
	if err != nil {
		log.Fatalf("Failed to create down migration file: %v", err)
	}
	defer downFile.Close()

	fmt.Printf("Created migration files:\n%s\n%s\n", upFileName, downFileName)
}

func runMigrations(direction string) {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		dbDriver := os.Getenv("DB_DRIVER")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_DATABASE")
		dbUser := os.Getenv("DB_USERNAME")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbSSLMode := os.Getenv("DB_SSLMODE")

		if dbDriver != "postgres" {
			log.Fatalf("Unsupported DB_DRIVER: %s", dbDriver)
		}

		dsn = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
			dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
	}

	// Create a new migration instance
	m, err := migrate.New(
		"file://migrations",
		dsn,
	)
	if err != nil {
		log.Fatal("Failed to create migration instance:", err)
	}

	// Run migrations based on the direction
	if direction == "up" {
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("Failed to apply migrations:", err)
		}
		log.Println("Migrations applied successfully")
	} else if direction == "down" {
		err = m.Down()
		if err != nil {
			log.Fatal("Failed to rollback migrations:", err)
		}
		log.Println("Migrations rolled back successfully")
	}
}
