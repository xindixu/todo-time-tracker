package db

import (
	"log"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// DBConnection holds the database connection and query builder
type DBConnection struct {
	DB      *sqlx.DB
	Builder *goqu.Database
}

// InitDatabaseConnection initializes the SQLite database connection (in-memory for now)
func InitDatabaseConnection() (*DBConnection, error) {
	log.Println("🔄 Initializing SQLite database (in-memory)...")

	// Open SQLite in-memory database
	db, err := sqlx.Connect("sqlite3", ":memory:")
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

	// Initialize goqu with SQLite dialect
	dialect := goqu.Dialect("sqlite3")
	builder := dialect.DB(db.DB)

	dbConnection := &DBConnection{
		DB:      db,
		Builder: builder,
	}

	// Run migrations
	if err := dbConnection.Migrate(); err != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error closing database connection:", err)
		}
		return nil, err
	}

	log.Println("✅ Database initialized successfully")
	return dbConnection, nil
}

// Migrate runs the database migrations
func (d *DBConnection) Migrate() error {
	log.Println("Running database migrations...")

	driver, err := sqlite3.WithInstance(d.DB.DB, &sqlite3.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"sqlite3", driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
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
