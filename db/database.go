package db

import (
	"log"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// DBConnection holds the database connection and query builder
type DBConnection struct {
	DB      *sqlx.DB
	Builder *goqu.Database
}

// InitDatabaseConnection initializes the PostgreSQL database connection
func InitDatabaseConnection(dbConnStr string) (*DBConnection, error) {
	log.Println("Initializing PostgreSQL database...")

	// Open PostgreSQL database
	db, err := sqlx.Connect("postgres", dbConnStr)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error closing database connection:", err)
		}
		return nil, err
	}

	// Initialize goqu with PostgreSQL dialect
	dialect := goqu.Dialect("postgres")
	builder := dialect.DB(db.DB)

	dbConnection := &DBConnection{
		DB:      db,
		Builder: builder,
	}

	log.Println("Database connection initialized successfully")

	return dbConnection, nil
}

// Migrate runs the database migrations
func (d *DBConnection) Migrate() error {
	log.Println("Running database migrations...")

	driver, err := postgres.WithInstance(d.DB.DB, &postgres.Config{})
	if err != nil {
		log.Printf("Failed to create postgres driver: %v", err)
		return err
	}

	// Use absolute path for migrations to avoid path resolution issues
	migrationPath := "file://./db/migrations"
	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres", driver,
	)
	if err != nil {
		log.Printf("Failed to create migration instance with path %s: %v", migrationPath, err)
		return err
	}

	// Check current version and dirty state
	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Printf("Failed to get migration version: %v", err)
		return err
	}

	if dirty {
		log.Printf("Database is in dirty state at version %d. Attempting to force version...", version)
		if err = m.Force(int(version)); err != nil {
			log.Printf("Failed to force version %d: %v", version, err)
			return err
		}
		log.Printf("Successfully forced version %d to clean state", version)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Migration failed: %v", err)
		return err
	}

	if err == migrate.ErrNoChange {
		log.Println("No new migrations to apply")
	} else {
		log.Println("Migrations applied successfully")
	}

	return nil
}

// Close closes the database connection
func (d *DBConnection) Close() error {
	log.Println("Closing database connection...")
	return d.DB.Close()
}

// Health checks if the database connection is healthy
func (d *DBConnection) Health() error {
	return d.DB.Ping()
}
