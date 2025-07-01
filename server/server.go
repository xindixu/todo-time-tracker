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
	utils "todo-time-tracker/go-utils"
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

// extractInfoFromRequest uses reflection to extract user info from any request with a Context field
// returns user uuid, user name, and token
func extractInfoFromRequest(req interface{}) (string, string, string) {
	if req == nil {
		return "", "", ""
	}

	// Use reflection to get the Context field
	reqValue := reflect.ValueOf(req)
	if reqValue.Kind() == reflect.Ptr {
		reqValue = reqValue.Elem()
	}

	// Look for a field named "Context"
	contextField := reqValue.FieldByName("Context")
	if !contextField.IsValid() || contextField.IsNil() {
		return "", "", ""
	}

	// Extract the context as *contextpb.Context
	if contextValue, ok := contextField.Interface().(*contextpb.Context); ok && contextValue != nil {
		return contextValue.UserUuid, contextValue.UserName, contextValue.Token
	}

	return "", "", ""
}

func populateContext(ctx context.Context, userUuid, userName, token string) context.Context {

	// Add username to the gRPC context for downstream use
	ctx = context.WithValue(ctx, utils.ContextKeyUserID, userUuid)
	ctx = context.WithValue(ctx, utils.ContextKeyToken, token)

	// TODO: find account id from user uuid
	// validate token, user uuid, user name

	ctx = context.WithValue(ctx, utils.ContextKeyAccountID, userUuid)

	return ctx
}

// RequestInterceptor is a unary server interceptor that logs requests and updates context
func RequestInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()

	// Extract request details for logging
	methodName := info.FullMethod
	userUuid, userName, token := extractInfoFromRequest(req)

	// Log incoming request
	log.Printf("[REQUEST] Method: %s, User: %s, Time: %s", methodName, userName, startTime.Format(time.RFC3339))

	// Validate that username is provided for all requests
	if userUuid == "" || userName == "" || token == "" {
		log.Printf("[ERROR] Request rejected - missing user info in context for method: %s", methodName)
		return nil, status.Error(codes.Unauthenticated, "user info is required in context")
	}

	// Add user info to the gRPC context for downstream use
	ctx = populateContext(ctx, userUuid, userName, token)
	ctx = context.WithValue(ctx, utils.ContextKeyRequestTime, startTime)

	// Call the actual handler
	resp, err := handler(ctx, req)

	// Calculate request duration
	duration := time.Since(startTime)

	// Log response
	if err != nil {
		log.Printf("[RESPONSE] Method: %s, User: %s, Duration: %v, Status: ERROR, Error: %v",
			methodName, userName, duration, err)
	} else {
		log.Printf("[RESPONSE] Method: %s, User: %s, Duration: %v, Status: SUCCESS",
			methodName, userName, duration)
	}

	return resp, err
}

// GetServerOptions returns gRPC server options with the interceptor
func GetServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(RequestInterceptor),
	}
}
