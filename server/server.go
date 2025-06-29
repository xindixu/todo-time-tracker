package server

import (
	"todo-time-tracker/proto/ttt"
)

// TTTServer implements the TTTService gRPC service
type TTTServer struct {
	ttt.UnimplementedTTTServiceServer
}

// NewTTTServer creates a new TTTServer instance
func NewTTTServer() *TTTServer {
	return &TTTServer{}
}
