package main

import (
	"context"
	"log"
	"os"

	"todo-time-tracker/db"
)

func main() {
	log.Println("🔄 Starting Todo Time Tracker database migration tool...")

	// Initialize database
	sqlDBConnStr := os.Getenv("SQL_DB_URL")
	graphDBConnArgs := db.GraphDBConnectionArgs{
		DBName:     os.Getenv("GRAPH_DB_NAME"),
		DBUri:      os.Getenv("GRAPH_DB_URI"),
		DBUser:     os.Getenv("GRAPH_DB_USER"),
		DBPassword: os.Getenv("GRAPH_DB_PASSWORD"),
	}
	dbConn, err := db.InitDatabaseConnection(context.Background(), sqlDBConnStr, graphDBConnArgs)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	err = dbConn.SQLDBMigrate()
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("✅ Database migrated successfully")

	defer func() {
		if err = dbConn.Close(context.Background()); err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}()
}
