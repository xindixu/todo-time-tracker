package server

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	utils "todo-time-tracker/go-utils"
	"todo-time-tracker/proto/go/model"
	"todo-time-tracker/proto/go/ttt"
)

// CreateTask creates a new task
func (s *TTTServer) CreateTask(ctx context.Context, req *ttt.CreateTaskReq) (*ttt.CreateTaskResp, error) {
	// Validate input (authentication is handled by interceptor)
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "task name cannot be empty")
	}

	// Validate status
	if req.Status == model.Task_NONE {
		return nil, status.Error(codes.InvalidArgument, "task status cannot be NONE")
	}

	// Generate UUID for the new task
	taskUUID := uuid.New()

	// Convert protobuf status to database status
	dbStatus, err := utils.TaskStatusPBToDB(req.Status)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, utils.WrapAsStr(err, "invalid task status"))
	}

	// Convert protobuf duration to time.Duration
	var estimatedDuration *time.Duration
	if req.EstimatedDuration != nil {
		duration := req.EstimatedDuration.AsDuration()
		estimatedDuration = &duration
	}

	// Create task in database
	dbTask, err := s.accessor.CreateTask(ctx, taskUUID, req.Name, req.Description, estimatedDuration, dbStatus)
	if err != nil {
		return nil, status.Error(codes.Internal, utils.WrapAsStr(err, "failed to create task"))
	}

	// Convert database model to protobuf model
	pbTask := &model.Task{
		Id:          dbTask.ID,
		Uuid:        dbTask.UUID.String(),
		Name:        dbTask.Name,
		Description: dbTask.Description,
		Status:      req.Status,
		CreatedAt:   timestamppb.New(dbTask.CreatedAt),
		UpdatedAt:   timestamppb.New(dbTask.UpdatedAt),
	}

	// Convert estimated duration back to protobuf duration
	if dbTask.EstimatedDuration != nil {
		pbTask.EstimatedDuration = durationpb.New(*dbTask.EstimatedDuration)
	}

	return &ttt.CreateTaskResp{
		Task: pbTask,
	}, nil
}

// GetTask retrieves a task by UUID
func (s *TTTServer) GetTask(ctx context.Context, req *ttt.GetTaskReq) (*ttt.GetTaskResp, error) {
	// Validate input
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "task UUID cannot be empty")
	}

	// Get task from database
	dbTask, err := s.accessor.GetTaskByUUID(ctx, req.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "task not found")
		}
		return nil, status.Error(codes.Internal, utils.WrapAsStr(err, "failed to get task"))
	}

	// Convert database model to protobuf model
	pbTask := &model.Task{
		Id:          dbTask.ID,
		Uuid:        dbTask.UUID.String(),
		Name:        dbTask.Name,
		Description: dbTask.Description,
		CreatedAt:   timestamppb.New(dbTask.CreatedAt),
		UpdatedAt:   timestamppb.New(dbTask.UpdatedAt),
	}

	// Convert status to protobuf enum
	pbStatus, err := utils.TaskStatusDBToPB(dbTask.Status)
	if err != nil {
		return nil, status.Error(codes.Internal, utils.WrapAsStr(err, "failed to convert task status"))
	}
	pbTask.Status = pbStatus

	// Convert estimated duration to protobuf duration
	if dbTask.EstimatedDuration != nil {
		pbTask.EstimatedDuration = durationpb.New(*dbTask.EstimatedDuration)
	}

	return &ttt.GetTaskResp{
		Task: pbTask,
	}, nil
}
