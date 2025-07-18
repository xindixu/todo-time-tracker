package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Connection holds the database connection and query builder
type Connection struct {
	SQLDB                 *sqlx.DB
	SQLBuilder            *goqu.Database
	GraphDB               neo4j.DriverWithContext
	GraphDBConnectionArgs GraphDBConnectionArgs
	SQLDBConnectionArgs   string
}

// findProjectRoot finds the project root directory by looking for go.mod file
func findProjectRoot() (string, error) {
	// Start from current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up the directory tree looking for go.mod
	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			// Reached root directory without finding go.mod
			return "", fmt.Errorf("could not find go.mod file in any parent directory")
		}
		currentDir = parent
	}
}

// InitDatabaseConnection initializes database connections
func InitDatabaseConnection(ctx context.Context, sqlDBConnStr string, graphDBConnArgs GraphDBConnectionArgs) (*Connection, error) {
	log.Println("Initializing databases...")

	sqlDB, sqlBuilder, err := InitSQLDBConnection(ctx, sqlDBConnStr)
	if err != nil {
		return nil, err
	}

	graphDB, err := InitGraphDBConnection(ctx, graphDBConnArgs)
	if err != nil {
		return nil, err
	}

	return &Connection{
		SQLDB:                 sqlDB,
		SQLBuilder:            sqlBuilder,
		GraphDB:               graphDB,
		GraphDBConnectionArgs: graphDBConnArgs,
		SQLDBConnectionArgs:   sqlDBConnStr,
	}, nil
}

// Close closes the database connection
func (d *Connection) Close(ctx context.Context) error {
	log.Println("Closing database connection...")
	err := d.SQLDB.Close()
	if err != nil {
		return err
	}
	err = d.GraphDB.Close(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Health checks if the database connection is healthy
func (d *Connection) Health(ctx context.Context) error {
	err := d.SQLDB.Ping()
	if err != nil {
		return err
	}
	err = d.GraphDB.VerifyConnectivity(ctx)
	if err != nil {
		return err
	}
	return nil
}
