package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"todo-time-tracker/db"
	"todo-time-tracker/db/accessors"
	"todo-time-tracker/proto/go/ttt"
	"todo-time-tracker/server"
)

const (
	port = ":50051"
)

func main() {
	log.Println("🔄 Starting Todo Time Tracker gRPC server...")

	ctx := context.Background()

	// Initialize databases
	sqlDBConnStr := os.Getenv("SQL_DB_URL")
	graphDBConnArgs := db.GraphDBConnectionArgs{
		DBUri:      os.Getenv("GRAPH_DB_URI"),
		DBUser:     os.Getenv("GRAPH_DB_USER"),
		DBPassword: os.Getenv("GRAPH_DB_PASSWORD"),
		DBName:     os.Getenv("GRAPH_DB_NAME"),
	}

	database, err := db.InitDatabaseConnection(ctx, sqlDBConnStr, graphDBConnArgs)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err = database.Close(ctx); err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}()

	// Create a TCP listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// Create accessor
	accessor := accessors.NewDBAccessor(database)

	// Create a new gRPC server with interceptors
	serverOptions := server.GetServerOptions(accessor)
	s := grpc.NewServer(serverOptions...)

	// Register the TTT service with database
	tttServer := server.NewTTTServer(database, accessor)
	ttt.RegisterTTTServiceServer(s, tttServer)

	// Register reflection service on gRPC server for development
	reflection.Register(s)

	log.Printf("gRPC server listening on port %s", port)
	log.Println("🚀 Server is ready to accept connections...")

	// Start serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
