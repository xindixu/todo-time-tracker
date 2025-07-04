package testhelpers_test

import (
	"database/sql"
	"fmt"
	"testing"
	"todo-time-tracker/db"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func createTestSQLDB(t *testing.T) (*sqlx.DB, *goqu.Database) {
	// Connect to default postgres database to create test database
	defaultConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		config.SQLDBUser, config.SQLDBPassword, config.SQLDBHost, config.SQLDBPort)

	defaultDB, err := sql.Open("postgres", defaultConnStr)
	require.NoError(t, err, "Failed to connect to default database")
	defer defaultDB.Close()

	// Create test database
	_, err = defaultDB.Exec(fmt.Sprintf("CREATE DATABASE %s", config.SQLDBName))
	require.NoError(t, err, "Failed to create test SQLDB")

	// Connect to test database
	sqlDB, sqlBuilder, err := db.InitSQLDBConnection(t.Context(), config.SQLDBConnStr)
	require.NoError(t, err, "Failed to initialize test database connection")

	return sqlDB, sqlBuilder
}

func dropTestSQLDB(t *testing.T, dbConnection *db.Connection) {
	// Connect to default postgres database to create test database
	defaultConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		config.SQLDBUser, config.SQLDBPassword, config.SQLDBHost, config.SQLDBPort)

	defaultDB, err := sql.Open("postgres", defaultConnStr)
	require.NoError(t, err, "Failed to connect to default database")
	defer defaultDB.Close()

	// Drop test database
	_, err = defaultDB.Exec(fmt.Sprintf("DROP DATABASE %s", config.SQLDBName))
	require.NoError(t, err, "Failed to drop test SQLDB")
}

// CleanupTestSQLDB removes all test data from the database
func CleanupTestSQLDB(t *testing.T, dbConnection *db.Connection) {
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
		_, err := dbConnection.SQLDB.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			t.Logf("Warning: Failed to cleanup table %s: %v", table, err)
		}
	}
}

// isPostgreSQLAvailable checks if PostgreSQL is available for testing
func isPostgreSQLAvailable(t *testing.T) bool {
	// Try to connect to default postgres database
	db, err := sql.Open("postgres", config.SQLDBConnStr)
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

// SkipIfNoPostgreSQL skips the test if PostgreSQL is not available
func SkipIfNoPostgreSQL(t *testing.T) {
	if !isPostgreSQLAvailable(t) {
		t.Skip("PostgreSQL not available, skipping test")
	}
}
