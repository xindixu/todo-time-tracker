package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"todo-time-tracker/db"
	"todo-time-tracker/proto/ttt"
	"todo-time-tracker/server"
)

const (
	port = ":50051"
)

func main() {
	log.Println("🔄 Starting Todo Time Tracker gRPC server...")

	// Initialize database
	dbConnStr := os.Getenv("DATABASE_URL")
	database, err := db.InitDatabaseConnection(dbConnStr)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err = database.DB.Close(); err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}()

	// Create a TCP listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// Create a new gRPC server with interceptors
	serverOptions := server.GetServerOptions()
	s := grpc.NewServer(serverOptions...)

	// Register the TTT service with database
	tttServer := server.NewTTTServer(database)
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
