package server

import (
	"context"
	"fmt"
	"time"
)

// generateUUID generates a simple UUID (replace with proper UUID library later)
func generateUUID() string {
	return fmt.Sprintf("tag-%d-%d", time.Now().UnixNano(), nextID)
}

// getUsername extracts username from the enhanced gRPC context
func getUsername(ctx context.Context) string {
	if username, ok := ctx.Value("username").(string); ok {
		return username
	}
	return "unknown" // This should not happen due to interceptor validation
}
