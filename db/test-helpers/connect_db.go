package testhelpers_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
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

// getEnvOrDefault returns environment variable value or default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// DefaultTestDBConfig returns default test database configuration
func DefaultTestDBConfig() *TestDBConfig {
	config := &TestDBConfig{
		SQLDBHost:     getEnvOrDefault("TEST_SQL_DB_HOST", "localhost"),
		SQLDBPort:     getEnvOrDefault("TEST_SQL_DB_PORT", "5432"),
		SQLDBUser:     getEnvOrDefault("TEST_SQL_DB_USER", "postgres"),
		SQLDBPassword: getEnvOrDefault("TEST_SQL_DB_PASSWORD", "postgres"),
		SQLDBName:     fmt.Sprintf("test_ttt_%d", time.Now().UnixNano()),

		GraphDBUri:      getEnvOrDefault("TEST_GRAPH_DB_URI", "neo4j://127.0.0.1:7687"),
		GraphDBUser:     getEnvOrDefault("TEST_GRAPH_DB_USER", "neo4j"),
		GraphDBPassword: getEnvOrDefault("TEST_GRAPH_DB_PASSWORD", "password"),
		GraphDBName:     fmt.Sprintf("test_ttt_%d", time.Now().UnixNano()),
	}

	config.SQLDBConnStr = fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		config.SQLDBUser, config.SQLDBPassword, config.SQLDBHost, config.SQLDBPort)

	return config
}

// isPostgreSQLAvailable checks if PostgreSQL is available for testing
func isPostgreSQLAvailable(t *testing.T) bool {
	config := DefaultTestDBConfig()

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

func createTestSQLDB(t *testing.T, config *TestDBConfig) (*sqlx.DB, *goqu.Database) {
	// Connect to default postgres database to create test database
	defaultConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		config.SQLDBUser, config.SQLDBPassword, config.SQLDBHost, config.SQLDBPort)

	defaultDB, err := sql.Open("postgres", defaultConnStr)
	require.NoError(t, err, "Failed to connect to default database")
	defer defaultDB.Close()

	// Drop test database if it exists
	_, err = defaultDB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", config.SQLDBName))
	require.NoError(t, err, "Failed to drop test database")

	// Create test database
	_, err = defaultDB.Exec(fmt.Sprintf("CREATE DATABASE %s", config.SQLDBName))
	require.NoError(t, err, "Failed to create test database")

	// Connect to test database
	sqlDB, sqlBuilder, err := db.InitSQLDBConnection(t.Context(), config.SQLDBConnStr)
	require.NoError(t, err, "Failed to initialize test database connection")

	return sqlDB, sqlBuilder
}

func createTestGraphDB(t *testing.T, config *TestDBConfig) neo4j.DriverWithContext {
	graphDB, err := db.InitGraphDBConnection(t.Context(), db.GraphDBConnectionArgs{
		DBUri:      config.GraphDBUri,
		DBUser:     config.GraphDBUser,
		DBPassword: config.GraphDBPassword,
		DBName:     config.GraphDBName,
	})
	if err != nil {
		t.Logf("Neo4j not available, skipping graph database tests: %v", err)
		return nil
	}

	return graphDB
}

// CreateTestDB creates a test database and returns connection
func CreateTestDB(t *testing.T, config *TestDBConfig) (*db.DBConnection, func()) {
	if config == nil {
		config = DefaultTestDBConfig()
	}

	sqlDB, sqlBuilder := createTestSQLDB(t, config)

	graphDB := createTestGraphDB(t, config)

	dbConnection := &db.DBConnection{
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
		err := dbConnection.Close(t.Context())
		if err != nil {
			t.Logf("Warning: Failed to close test database for cleanup: %v", err)
		}

		// // Reconnect to default database to drop test database
		// defaultDB, err := sql.Open("postgres", defaultConnStr)
		// if err != nil {
		// 	t.Logf("Warning: Failed to connect to default database for cleanup: %v", err)
		// 	return
		// }
		// err = defaultDB.Close()
		// if err != nil {
		// 	t.Logf("Warning: Failed to close default database for cleanup: %v", err)
		// }

		// // Terminate all connections to the test database
		// _, err = defaultDB.Exec(fmt.Sprintf(`
		// 	SELECT pg_terminate_backend(pid)
		// 	FROM pg_stat_activity
		// 	WHERE datname = '%s' AND pid <> pg_backend_pid()
		// `, config.SQLDBName))
		// if err != nil {
		// 	t.Logf("Warning: Failed to terminate connections: %v", err)
		// }

		// // Drop test database
		// _, err = defaultDB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", config.SQLDBName))
		// if err != nil {
		// 	t.Logf("Warning: Failed to drop test database: %v", err)
		// }
	}

	return dbConnection, cleanup
}

// CleanupTestSQLDB removes all test data from the database
func CleanupTestSQLDB(t *testing.T, dbConnection *db.DBConnection) {
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

// SkipIfNoPostgreSQL skips the test if PostgreSQL is not available
func SkipIfNoPostgreSQL(t *testing.T) {
	if !isPostgreSQLAvailable(t) {
		t.Skip("PostgreSQL not available, skipping test")
	}
}

// isNeo4jAvailable checks if Neo4j is available for testing
func isNeo4jAvailable(t *testing.T) bool {
	config := DefaultTestDBConfig()

	// Try to connect to Neo4j
	driver, err := neo4j.NewDriverWithContext(config.GraphDBUri, neo4j.BasicAuth(config.GraphDBUser, config.GraphDBPassword, ""))
	if err != nil {
		t.Logf("Neo4j not available: %v", err)
		return false
	}
	defer driver.Close(t.Context())

	// Test connection
	if err := driver.VerifyConnectivity(t.Context()); err != nil {
		t.Logf("Neo4j not available: %v", err)
		return false
	}

	return true
}

// SkipIfNoNeo4j skips the test if Neo4j is not available
func SkipIfNoNeo4j(t *testing.T) {
	if !isNeo4jAvailable(t) {
		t.Skip("Neo4j not available, skipping test")
	}
}

func CreateTestDBAccessor(t *testing.T, dbConnection *db.DBConnection) *TestDBAccessor {
	return NewTestDBAccessor(dbConnection)
}
