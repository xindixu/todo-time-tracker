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

type createTaskReqValidator struct {
	Name              string               `validate:"required"`
	Description       string               `validate:"omitempty,max=255"`
	Status            model.Task_Status    `validate:"oneof=TODO IN_PROGRESS DONE BLOCKED"`
	EstimatedDuration *durationpb.Duration `validate:"omitempty,duration,gt=0"`
}

// CreateTask creates a new task
func (s *TTTServer) CreateTask(ctx context.Context, req *ttt.CreateTaskReq) (*ttt.CreateTaskResp, error) {
	// Validate input (authentication is handled by interceptor)
	validator := createTaskReqValidator{
		Name:              req.Name,
		Description:       req.Description,
		Status:            req.Status,
		EstimatedDuration: req.EstimatedDuration,
	}
	err := validate.Struct(validator)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, utils.WrapAsStr(err, "invalid request"))
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

type getTaskReqValidator struct {
	UUID string `validate:"required,uuid"`
}

// GetTask retrieves a task by UUID
func (s *TTTServer) GetTask(ctx context.Context, req *ttt.GetTaskReq) (*ttt.GetTaskResp, error) {
	// Validate input
	validator := getTaskReqValidator{
		UUID: req.Uuid,
	}
	err := validate.Struct(validator)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, utils.WrapAsStr(err, "invalid request"))
	}

	taskUUID := uuid.MustParse(req.Uuid)
	// Get task from database
	dbTask, err := s.accessor.GetTaskByUUID(ctx, taskUUID)
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

type createTaskLinksReqValidator struct {
	FromTaskUUID string          `validate:"required,uuid"`
	ToTaskUUID   string          `validate:"required,uuid"`
	Link         model.Task_Link `validate:"oneof=PARENT_OF BLOCKS RELATES_TO DUPLICATE_OF"`
}

// CreateTaskLinks links two tasks together
func (s *TTTServer) CreateTaskLinks(ctx context.Context, req *ttt.CreateTaskLinksReq) (*ttt.CreateTaskLinksResp, error) {
	// Validate input
	validator := createTaskLinksReqValidator{
		FromTaskUUID: req.FromTaskUuid,
		ToTaskUUID:   req.ToTaskUuid,
		Link:         req.Link,
	}
	err := validate.Struct(validator)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, utils.WrapAsStr(err, "invalid request"))
	}

	// Convert protobuf link to database link
	dbLink, err := utils.TaskLinkPBToDB(req.Link)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, utils.WrapAsStr(err, "invalid task link"))
	}

	// Convert protobuf UUIDs to database UUIDs
	fromTaskUUID := uuid.MustParse(req.FromTaskUuid)
	toTaskUUID := uuid.MustParse(req.ToTaskUuid)

	// Link tasks
	err = s.accessor.CreateTaskLinks(ctx, fromTaskUUID, toTaskUUID, dbLink)
	if err != nil {
		return nil, status.Error(codes.Internal, utils.WrapAsStr(err, "failed to link tasks"))
	}

	return &ttt.CreateTaskLinksResp{}, nil
}

type getTaskLinksReqValidator struct {
	FromTaskUUID string `validate:"required,uuid"`
	ToTaskUUID   string `validate:"required,uuid"`
}

// GetTaskLinks retrieves links between two tasks
func (s *TTTServer) GetTaskLinks(ctx context.Context, req *ttt.GetTaskLinksReq) (*ttt.GetTaskLinksResp, error) {
	// Validate input
	validator := getTaskLinksReqValidator{
		FromTaskUUID: req.FromTaskUuid,
		ToTaskUUID:   req.ToTaskUuid,
	}
	err := validate.Struct(validator)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, utils.WrapAsStr(err, "invalid request"))
	}

	fromTaskUUID := uuid.MustParse(req.FromTaskUuid)
	toTaskUUID := uuid.MustParse(req.ToTaskUuid)

	links, err := s.accessor.GetTaskLinks(ctx, fromTaskUUID, toTaskUUID)
	if err != nil {
		return nil, status.Error(codes.Internal, utils.WrapAsStr(err, "failed to get task links"))
	}

	pbLinks := make([]model.Task_Link, 0)
	for _, link := range links {
		pbLink, err := utils.TaskLinkDBToPB(link)
		if err != nil {
			return nil, status.Error(codes.Internal, utils.WrapAsStr(err, "failed to convert task link"))
		}
		pbLinks = append(pbLinks, pbLink)
	}

	return &ttt.GetTaskLinksResp{Links: pbLinks}, nil
}
