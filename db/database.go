package db

import (
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

// DBConnection holds the database connection and query builder
type DBConnection struct {
	DB      *sql.DB
	Builder *goqu.Database
}

// InitDatabase initializes the SQLite database connection (in-memory for now)
func InitDatabaseConnection() (*DBConnection, error) {
	log.Println("Initializing SQLite database (in-memory)...")

	// Open SQLite in-memory database
	db, err := sql.Open("sqlite3", ":memory:")
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
	builder := goqu.New("sqlite3", db)

	dbConnection := &DBConnection{
		DB:      db,
		Builder: builder,
	}

	// Create tables
	if err := dbConnection.createTables(); err != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error closing database connection:", err)
		}
		return nil, err
	}

	log.Println("✅ Database initialized successfully")
	return dbConnection, nil
}

// createTables creates the necessary database tables
func (d *DBConnection) createTables() error {
	log.Println("Creating database tables...")

	// Create tags table
	createTagsTable := `
		CREATE TABLE IF NOT EXISTS tags (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX IF NOT EXISTS idx_tags_uuid ON tags(uuid);
		CREATE INDEX IF NOT EXISTS idx_tags_name ON tags(name);
	`

	if _, err := d.DB.Exec(createTagsTable); err != nil {
		return err
	}

	log.Println("✅ Tables created successfully")
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
