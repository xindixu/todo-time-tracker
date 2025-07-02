package server

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	utils "todo-time-tracker/go-utils"
	"todo-time-tracker/proto/go/model"
	"todo-time-tracker/proto/go/ttt"
)

func (s *TTTServer) CreateUser(ctx context.Context, req *ttt.CreateUserReq) (*ttt.CreateUserResp, error) {

	// Validate input
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}

	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email cannot be empty")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password cannot be empty")
	}

	// Generate UUID for the new user
	uuid := uuid.New()

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, utils.WrapAsStr(err, "failed to hash password"))
	}

	// Create user in database
	dbUser, err := s.accessor.CreateUser(ctx, uuid, req.Name, req.Email, string(hashedPassword))
	if err != nil {
		return nil, status.Error(codes.Internal, utils.WrapAsStr(err, "failed to create user"))
	}

	// Convert database model to protobuf model
	protoUser := &model.User{
		Id:    dbUser.ID,
		Uuid:  dbUser.UUID.String(),
		Name:  dbUser.Name,
		Email: dbUser.Email,
	}

	return &ttt.CreateUserResp{
		User: protoUser,
	}, nil
}
