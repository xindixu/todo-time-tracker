package main

import (
	"log"
	"os"

	"todo-time-tracker/db"
)

func main() {
	log.Println("🔄 Starting Todo Time Tracker database migration tool...")

	// Initialize database
	dbConnStr := os.Getenv("DATABASE_URL")
	database, err := db.InitDatabaseConnection(dbConnStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	err = database.Migrate()
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("✅ Database migrated successfully")

	defer func() {
		if err = database.DB.Close(); err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}()
}
