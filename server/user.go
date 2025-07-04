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

type createUserReqValidator struct {
	Name     string `validate:"required,min=3,max=255"`
	Email    string `validate:"required,email,min=1,max=255"`
	Password string `validate:"required,min=8,max=255"`
}

// CreateUser creates a new user
func (s *TTTServer) CreateUser(ctx context.Context, req *ttt.CreateUserReq) (*ttt.CreateUserResp, error) {
	validator := createUserReqValidator{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err := validate.Struct(validator)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, utils.WrapAsStr(err, "invalid request"))
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
