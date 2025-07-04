package testhelpers_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"todo-time-tracker/db"
)

// TestDBConfig holds configuration for test database
type TestDBConfig struct {
	SQLDBHost     string
	SQLDBPort     string
	SQLDBUser     string
	SQLDBPassword string
	SQLDBName     string
	SQLDBConnStr  string

	GraphDBUri      string
	GraphDBUser     string
	GraphDBPassword string
	GraphDBName     string
}

var config = getDefaultTestDBConfig()

// getEnvOrDefault returns environment variable value or default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// DefaultTestDBConfig returns default test database configuration
func getDefaultTestDBConfig() *TestDBConfig {
	config := &TestDBConfig{
		SQLDBHost:     getEnvOrDefault("TEST_SQL_DB_HOST", "localhost"),
		SQLDBPort:     getEnvOrDefault("TEST_SQL_DB_PORT", "5432"),
		SQLDBUser:     getEnvOrDefault("TEST_SQL_DB_USER", "postgres"),
		SQLDBPassword: getEnvOrDefault("TEST_SQL_DB_PASSWORD", "postgres"),
		SQLDBName:     fmt.Sprintf("test_ttt_%d", time.Now().UnixNano()),

		GraphDBUri:      getEnvOrDefault("TEST_GRAPH_DB_URI", "neo4j://127.0.0.1:7687"),
		GraphDBUser:     getEnvOrDefault("TEST_GRAPH_DB_USER", "neo4j"),
		GraphDBPassword: getEnvOrDefault("TEST_GRAPH_DB_PASSWORD", "password"),
		GraphDBName:     fmt.Sprintf("test-ttt-%d", time.Now().UnixNano()),
	}

	config.SQLDBConnStr = fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		config.SQLDBUser, config.SQLDBPassword, config.SQLDBHost, config.SQLDBPort)

	return config
}

// CreateTestDB creates a test database and returns connection
func CreateTestDB(t *testing.T) (*db.Connection, func()) {
	sqlDB, sqlBuilder := createTestSQLDB(t)

	graphDB := createTestGraphDB(t)

	dbConnection := &db.Connection{
		SQLDB:      sqlDB,
		SQLBuilder: sqlBuilder,
		GraphDB:    graphDB,
		GraphDBConnectionArgs: db.GraphDBConnectionArgs{
			DBUri:      config.GraphDBUri,
			DBUser:     config.GraphDBUser,
			DBPassword: config.GraphDBPassword,
			DBName:     config.GraphDBName,
		},
	}

	// Run migrations
	err := dbConnection.SQLDBMigrate()
	require.NoError(t, err, "Failed to run database migrations")

	// Return cleanup function
	cleanup := func() {
		log.Println("Cleaning up test database")

		// Drop databases before closing the connection
		dropTestSQLDB(t, dbConnection)
		dropTestGraphDB(t, dbConnection)

		// Close the connection last
		err := dbConnection.Close(t.Context())
		if err != nil {
			t.Logf("Warning: Failed to close test database for cleanup: %v", err)
		}
	}

	return dbConnection, cleanup
}

// CreateTestDBAccessor creates a test database accessor
func CreateTestDBAccessor(t *testing.T, dbConnection *db.Connection) *TestDBAccessor {
	return NewTestDBAccessor(dbConnection)
}
