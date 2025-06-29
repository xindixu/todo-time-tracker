package server

import (
	"context"

	"github.com/google/uuid"
)

// generateUUID generates a standard UUID using Google's UUID package
func generateUUID() string {
	return uuid.New().String()
}

// getUsername extracts username from the enhanced gRPC context
func getUsername(ctx context.Context) string {
	if username, ok := ctx.Value("username").(string); ok {
		return username
	}
	return "unknown" // This should not happen due to interceptor validation
}
