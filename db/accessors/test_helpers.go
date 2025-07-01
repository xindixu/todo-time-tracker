package accessors

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"todo-time-tracker/db"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

// TestDBConfig holds configuration for test database
type TestDBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// DefaultTestDBConfig returns default test database configuration
func DefaultTestDBConfig() *TestDBConfig {
	return &TestDBConfig{
		Host:     getEnvOrDefault("TEST_DB_HOST", "localhost"),
		Port:     getEnvOrDefault("TEST_DB_PORT", "5432"),
		User:     getEnvOrDefault("TEST_DB_USER", "postgres"),
		Password: getEnvOrDefault("TEST_DB_PASSWORD", "postgres"),
		DBName:   fmt.Sprintf("test_ttt_%d", time.Now().UnixNano()),
	}
}

// getEnvOrDefault returns environment variable value or default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// IsPostgreSQLAvailable checks if PostgreSQL is available for testing
func IsPostgreSQLAvailable(t *testing.T) bool {
	config := DefaultTestDBConfig()

	// Try to connect to default postgres database
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		config.User, config.Password, config.Host, config.Port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Logf("PostgreSQL not available: %v", err)
		return false
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		t.Logf("PostgreSQL not available: %v", err)
		return false
	}

	return true
}

// CreateTestDatabase creates a test database and returns connection
func CreateTestDatabase(t *testing.T, config *TestDBConfig) (*db.DBConnection, func()) {
	if config == nil {
		config = DefaultTestDBConfig()
	}

	// Connect to default postgres database to create test database
	defaultConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		config.User, config.Password, config.Host, config.Port)

	defaultDB, err := sql.Open("postgres", defaultConnStr)
	require.NoError(t, err, "Failed to connect to default database")
	defer defaultDB.Close()

	// Create test database
	_, err = defaultDB.Exec(fmt.Sprintf("CREATE DATABASE %s", config.DBName))
	require.NoError(t, err, "Failed to create test database")

	// Connect to test database
	testConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DBName)

	dbConnection, err := db.InitDatabaseConnection(testConnStr)
	require.NoError(t, err, "Failed to initialize test database connection")

	// Run migrations
	err = dbConnection.Migrate()
	require.NoError(t, err, "Failed to run database migrations")

	// Return cleanup function
	cleanup := func() {
		log.Println("Cleaning up test database")
		dbConnection.Close()

		// Reconnect to default database to drop test database
		defaultDB, err := sql.Open("postgres", defaultConnStr)
		if err != nil {
			t.Logf("Warning: Failed to connect to default database for cleanup: %v", err)
			return
		}
		defer defaultDB.Close()

		// Terminate all connections to the test database
		_, err = defaultDB.Exec(fmt.Sprintf(`
			SELECT pg_terminate_backend(pid)
			FROM pg_stat_activity
			WHERE datname = '%s' AND pid <> pg_backend_pid()
		`, config.DBName))
		if err != nil {
			t.Logf("Warning: Failed to terminate connections: %v", err)
		}

		// Drop test database
		_, err = defaultDB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", config.DBName))
		if err != nil {
			t.Logf("Warning: Failed to drop test database: %v", err)
		}
	}

	return dbConnection, cleanup
}

// CleanupTestData removes all test data from the database
func CleanupTestData(t *testing.T, dbConnection *db.DBConnection) {
	// Delete in reverse order of dependencies
	tables := []string{
		"task_tags",
		"task_users",
		"organization_users",
		"sessions",
		"tasks",
		"tags",
		"users",
		"organizations",
		"accounts",
	}

	for _, table := range tables {
		_, err := dbConnection.DB.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			t.Logf("Warning: Failed to cleanup table %s: %v", table, err)
		}
	}
}

// SkipIfNoPostgreSQL skips the test if PostgreSQL is not available
func SkipIfNoPostgreSQL(t *testing.T) {
	if !IsPostgreSQLAvailable(t) {
		t.Skip("PostgreSQL not available, skipping test")
	}
}
