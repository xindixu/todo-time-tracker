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
	log.Println("Initializing graph database connection...")

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
