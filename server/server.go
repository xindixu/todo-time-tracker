package server

import (
	"context"
	"log"
	"reflect"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"todo-time-tracker/db"
	"todo-time-tracker/db/accessors"
	contextpb "todo-time-tracker/proto/go/context"
	"todo-time-tracker/proto/go/ttt"
)

// TTTServer implements the TTTService gRPC service
type TTTServer struct {
	ttt.UnimplementedTTTServiceServer
	db       *db.DBConnection
	accessor *accessors.DBAccessor
}

// NewTTTServer creates a new TTTServer instance with database connection
func NewTTTServer(db *db.DBConnection) *TTTServer {
	return &TTTServer{
		db:       db,
		accessor: accessors.NewDBAccessor(db),
	}
}

// extractUsernameFromRequest uses reflection to extract username from any request with a Context field
func extractUsernameFromRequest(req interface{}) string {
	if req == nil {
		return ""
	}

	// Use reflection to get the Context field
	reqValue := reflect.ValueOf(req)
	if reqValue.Kind() == reflect.Ptr {
		reqValue = reqValue.Elem()
	}

	// Look for a field named "Context"
	contextField := reqValue.FieldByName("Context")
	if !contextField.IsValid() || contextField.IsNil() {
		return ""
	}

	// Extract the context as *contextpb.Context
	if contextValue, ok := contextField.Interface().(*contextpb.Context); ok && contextValue != nil {
		return contextValue.Username
	}

	return ""
}

// RequestInterceptor is a unary server interceptor that logs requests and updates context
func RequestInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()

	// Extract request details for logging
	methodName := info.FullMethod
	username := extractUsernameFromRequest(req)

	// Log incoming request
	log.Printf("[REQUEST] Method: %s, User: %s, Time: %s", methodName, username, startTime.Format(time.RFC3339))

	// Validate that username is provided for all requests
	if username == "" {
		log.Printf("[ERROR] Request rejected - missing username in context for method: %s", methodName)
		return nil, status.Error(codes.Unauthenticated, "username is required in context")
	}

	// Add username to the gRPC context for downstream use
	ctx = context.WithValue(ctx, "username", username)
	ctx = context.WithValue(ctx, "request_start_time", startTime)

	// Call the actual handler
	resp, err := handler(ctx, req)

	// Calculate request duration
	duration := time.Since(startTime)

	// Log response
	if err != nil {
		log.Printf("[RESPONSE] Method: %s, User: %s, Duration: %v, Status: ERROR, Error: %v",
			methodName, username, duration, err)
	} else {
		log.Printf("[RESPONSE] Method: %s, User: %s, Duration: %v, Status: SUCCESS",
			methodName, username, duration)
	}

	return resp, err
}

// GetServerOptions returns gRPC server options with the interceptor
func GetServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(RequestInterceptor),
	}
}
