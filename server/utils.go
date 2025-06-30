package server

import (
	"context"
)

// getUsername extracts username from the enhanced gRPC context
func getUsername(ctx context.Context) string {
	if username, ok := ctx.Value("username").(string); ok {
		return username
	}
	return "unknown" // This should not happen due to interceptor validation
}
