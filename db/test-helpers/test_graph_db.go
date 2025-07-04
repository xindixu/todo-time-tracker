package testhelpers_test

import (
	"fmt"
	"testing"
	"todo-time-tracker/db"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/require"
)

func createTestGraphDB(t *testing.T) neo4j.DriverWithContext {
	graphDB, err := db.InitGraphDBConnection(t.Context(), db.GraphDBConnectionArgs{
		DBUri:      config.GraphDBUri,
		DBUser:     config.GraphDBUser,
		DBPassword: config.GraphDBPassword,
		DBName:     config.GraphDBName,
	})
	require.NoError(t, err, "Failed to connect to Neo4j")

	err = graphDB.VerifyConnectivity(t.Context())
	require.NoError(t, err, "Failed to verify connectivity to Neo4j")

	// Connect to default database to create test database
	session := graphDB.NewSession(t.Context(), neo4j.SessionConfig{
		DatabaseName: "system",
		AccessMode:   neo4j.AccessModeWrite,
	})

	// Create test database
	_, err = session.ExecuteWrite(t.Context(), func(tx neo4j.ManagedTransaction) (any, error) {
		cypher := fmt.Sprintf("CREATE DATABASE `%s` IF NOT EXISTS;", config.GraphDBName)
		_, err := tx.Run(t.Context(), cypher, nil)
		return nil, err
	})
	require.NoError(t, err, "Failed to create test GraphDB")

	return graphDB
}

func dropTestGraphDB(t *testing.T, dbConnection *db.Connection) {
	// Check if the driver is still open before trying to create a session
	// if dbConnection.GraphDB == nil {
	// 	t.Log("GraphDB driver is nil, skipping drop operation")
	// 	return
	// }

	// Try to create a session, but handle the case where the driver is closed
	session := dbConnection.GraphDB.NewSession(t.Context(), neo4j.SessionConfig{
		DatabaseName: "system",
		AccessMode:   neo4j.AccessModeWrite,
	})
	defer session.Close(t.Context())

	_, err := session.ExecuteWrite(t.Context(), func(tx neo4j.ManagedTransaction) (any, error) {
		cypher := fmt.Sprintf("DROP DATABASE `%s` IF EXISTS;", config.GraphDBName)
		_, err := tx.Run(t.Context(), cypher, nil)
		return nil, err
	})
	require.NoError(t, err, "Failed to drop GraphDB")
}

// isNeo4jAvailable checks if Neo4j is available for testing
func isNeo4jAvailable(t *testing.T) bool {
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
