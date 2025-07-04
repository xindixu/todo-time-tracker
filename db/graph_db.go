package db

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// GraphDBConnectionArgs contains the arguments for initializing the Neo4j database connection
type GraphDBConnectionArgs struct {
	DBUri      string
	DBUser     string
	DBPassword string
	DBName     string
}

// InitGraphDBConnection initializes the Neo4j database connection
func InitGraphDBConnection(ctx context.Context, args GraphDBConnectionArgs) (neo4j.DriverWithContext, error) {
	log.Println("Initializing Neo4j DB connection...")

	driver, err := neo4j.NewDriverWithContext(args.DBUri, neo4j.BasicAuth(args.DBUser, args.DBPassword, ""))
	if err != nil {
		return nil, err
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (dbConnection *DBConnection) NewGraphDBSession(ctx context.Context) (neo4j.SessionWithContext, error) {
	session := dbConnection.GraphDB.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: dbConnection.GraphDBConnectionArgs.DBName,
		AccessMode:   neo4j.AccessModeWrite,
	})

	return session, nil
}
