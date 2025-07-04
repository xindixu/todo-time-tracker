package server

import (
	"context"
	goutils "todo-time-tracker/go-utils"

	"github.com/go-playground/validator/v10"
)

// getUsername extracts username from the enhanced gRPC context
func getUsername(ctx context.Context) string {
	if username, ok := ctx.Value(goutils.ContextKeyUserName).(string); ok {
		return username
	}
	return "unknown" // This should not happen due to interceptor validation
}

var validate = validator.New()
