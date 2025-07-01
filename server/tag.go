package server

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	utils "todo-time-tracker/go-utils"
	"todo-time-tracker/proto/go/model"
	"todo-time-tracker/proto/go/ttt"
)

// CreateTag creates a new tag
func (s *TTTServer) CreateTag(ctx context.Context, req *ttt.CreateTagReq) (*ttt.CreateTagResp, error) {
	username := getUsername(ctx)
	log.Printf("CreateTag: Processing request for user: %s, name: %s", username, req.Name)

	// Validate input (authentication is handled by interceptor)
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "tag name cannot be empty")
	}

	// Generate UUID for the new tag
	uuid := uuid.New()

	// Create tag in database
	dbTag, err := s.accessor.CreateTag(ctx, uuid, req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, utils.WrapAsStr(err, "failed to create tag"))
	}

	// Convert database model to protobuf model
	protoTag := &model.Tag{
		Id:        dbTag.ID,
		Uuid:      dbTag.UUID.String(),
		Name:      dbTag.Name,
		CreatedAt: timestamppb.New(dbTag.CreatedAt),
		UpdatedAt: timestamppb.New(dbTag.UpdatedAt),
	}

	log.Printf("CreateTag: Successfully created tag: %s (UUID: %s) by user: %s", protoTag.Name, protoTag.Uuid, username)

	return &ttt.CreateTagResp{
		Tag: protoTag,
	}, nil
}

// GetTag retrieves a tag by UUID
func (s *TTTServer) GetTag(ctx context.Context, req *ttt.GetTagReq) (*ttt.GetTagResp, error) {
	username := getUsername(ctx)
	log.Printf("GetTag: Processing request for user: %s, UUID: %s", username, req.Uuid)

	// Validate input
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "tag UUID cannot be empty")
	}

	// Get tag from database
	dbTag, err := s.accessor.GetTagByUUID(ctx, req.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "tag not found")
		}
		return nil, status.Error(codes.Internal, utils.WrapAsStr(err, "failed to get tag"))
	}

	// Convert database model to protobuf model
	protoTag := &model.Tag{
		Id:        dbTag.ID,
		Uuid:      dbTag.UUID.String(),
		Name:      dbTag.Name,
		CreatedAt: timestamppb.New(dbTag.CreatedAt),
		UpdatedAt: timestamppb.New(dbTag.UpdatedAt),
	}

	log.Printf("GetTag: Successfully retrieved tag: %s (UUID: %s) for user: %s", protoTag.Name, protoTag.Uuid, username)

	return &ttt.GetTagResp{
		Tag: protoTag,
	}, nil
}

// // ListTags lists all tags
// func (s *TTTServer) ListTags(ctx context.Context, req *ttt.ListTagsReq) (*ttt.ListTagsResp, error) {
// 	username := getUsername(ctx)
// 	log.Printf("ListTags: Processing request for user: %s", username)

// 	// Get tags from database (filter by creator)
// 	dbTags, err := s.tagRepo.List(ctx, username)
// 	if err != nil {
// 		log.Printf("ListTags: Database error: %v", err)
// 		return nil, status.Error(codes.Internal, "failed to list tags")
// 	}

// 	// Convert database models to protobuf models
// 	var protoTags []*model.Tag
// 	for _, dbTag := range dbTags {
// 		protoTag := &model.Tag{
// 			Id:   dbTag.ID,
// 			Uuid: dbTag.UUID,
// 			Name: dbTag.Name,
// 		}
// 		protoTags = append(protoTags, protoTag)
// 	}

// 	log.Printf("ListTags: Successfully listed %d tags for user: %s", len(protoTags), username)

// 	return &ttt.ListTagsResp{
// 		Tags: protoTags,
// 	}, nil
// }

// // UpdateTag updates an existing tag
// func (s *TTTServer) UpdateTag(ctx context.Context, req *ttt.UpdateTagReq) (*ttt.UpdateTagResp, error) {
// 	username := getUsername(ctx)
// 	log.Printf("UpdateTag: Processing request for user: %s, UUID: %s, new name: %s", username, req.Uuid, req.Name)

// 	// Validate input
// 	if req.Uuid == "" {
// 		return nil, status.Error(codes.InvalidArgument, "tag UUID cannot be empty")
// 	}
// 	if req.Name == "" {
// 		return nil, status.Error(codes.InvalidArgument, "tag name cannot be empty")
// 	}

// 	// Update tag in database
// 	dbTag, err := s.tagRepo.Update(ctx, req.Uuid, req.Name)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, status.Error(codes.NotFound, "tag not found")
// 		}
// 		log.Printf("UpdateTag: Database error: %v", err)
// 		return nil, status.Error(codes.Internal, "failed to update tag")
// 	}

// 	// Convert database model to protobuf model
// 	protoTag := &model.Tag{
// 		Id:   dbTag.ID,
// 		Uuid: dbTag.UUID,
// 		Name: dbTag.Name,
// 	}

// 	log.Printf("UpdateTag: Successfully updated tag: %s (UUID: %s) by user: %s", protoTag.Name, protoTag.Uuid, username)

// 	return &ttt.UpdateTagResp{
// 		Tag: protoTag,
// 	}, nil
// }

// // DeleteTag deletes a tag by UUID
// func (s *TTTServer) DeleteTag(ctx context.Context, req *ttt.DeleteTagReq) (*ttt.DeleteTagResp, error) {
// 	username := getUsername(ctx)
// 	log.Printf("DeleteTag: Processing request for user: %s, UUID: %s", username, req.Uuid)

// 	// Validate input
// 	if req.Uuid == "" {
// 		return nil, status.Error(codes.InvalidArgument, "tag UUID cannot be empty")
// 	}

// 	// Check if tag exists first
// 	exists, err := s.tagRepo.Exists(ctx, req.Uuid)
// 	if err != nil {
// 		log.Printf("DeleteTag: Database error checking existence: %v", err)
// 		return nil, status.Error(codes.Internal, "failed to check tag existence")
// 	}
// 	if !exists {
// 		return nil, status.Error(codes.NotFound, "tag not found")
// 	}

// 	// Delete tag from database
// 	err = s.tagRepo.Delete(ctx, req.Uuid)
// 	if err != nil {
// 		log.Printf("DeleteTag: Database error: %v", err)
// 		return nil, status.Error(codes.Internal, "failed to delete tag")
// 	}

// 	log.Printf("DeleteTag: Successfully deleted tag (UUID: %s) by user: %s", req.Uuid, username)

// 	return &ttt.DeleteTagResp{}, nil
// }
