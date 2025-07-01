package rpc

import (
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"todo-time-tracker/proto/go/ttt"
)

const (
	defaultServerAddress = "localhost:50051"
	defaultTimeout       = 10 * time.Second
)

// TTTClient wraps the gRPC client with convenience methods
type TTTClient struct {
	client   ttt.TTTServiceClient
	conn     *grpc.ClientConn
	username string
}

// NewTTTClient creates a new TTT client connection
func NewTTTClient(serverAddress, username string) (*TTTClient, error) {
	if serverAddress == "" {
		serverAddress = defaultServerAddress
	}

	if username == "" {
		// Try to get username from environment or use a default
		username = os.Getenv("USER")
		if username == "" {
			username = "cli-user"
		}
	}

	// Connect to the gRPC server
	conn, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}

	client := ttt.NewTTTServiceClient(conn)

	return &TTTClient{
		client:   client,
		conn:     conn,
		username: username,
	}, nil
}

// Close closes the gRPC connection
func (c *TTTClient) Close() error {
	return c.conn.Close()
}

// logGRPCError logs gRPC errors in a user-friendly way
func logGRPCError(err error) {
	log.Printf("gRPC error: %v", err)
}
