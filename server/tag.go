package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"todo-time-tracker/proto/tag"
	"todo-time-tracker/proto/ttt"
)

// In-memory storage for tags (replace with database later)
var (
	tags   = make(map[string]*tag.Tag)
	tagsMu sync.RWMutex
	nextID int64 = 1
)

// generateUUID generates a simple UUID (replace with proper UUID library later)
func generateUUID() string {
	return fmt.Sprintf("tag-%d-%d", time.Now().UnixNano(), nextID)
}

// CreateTag creates a new tag
func (s *TTTServer) CreateTag(ctx context.Context, req *ttt.CreateTagReq) (*ttt.CreateTagResp, error) {
	log.Printf("CreateTag called by user: %s, name: %s", req.Context.Username, req.Name)

	// Validate input
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "tag name cannot be empty")
	}

	// Validate context
	if req.Context == nil || req.Context.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "context with username is required")
	}

	tagsMu.Lock()
	defer tagsMu.Unlock()

	// Create new tag
	newTag := &tag.Tag{
		Id:   nextID,
		Uuid: generateUUID(),
		Name: req.Name,
	}

	nextID++
	tags[newTag.Uuid] = newTag

	log.Printf("Created tag: %s (UUID: %s) by user: %s", newTag.Name, newTag.Uuid, req.Context.Username)

	return &ttt.CreateTagResp{
		Tag: newTag,
	}, nil
}

// GetTag retrieves a tag by UUID
func (s *TTTServer) GetTag(ctx context.Context, req *ttt.GetTagReq) (*ttt.GetTagResp, error) {
	log.Printf("GetTag called by user: %s, UUID: %s", req.Context.Username, req.Uuid)

	// Validate input
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "tag UUID cannot be empty")
	}

	// Validate context
	if req.Context == nil || req.Context.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "context with username is required")
	}

	tagsMu.RLock()
	defer tagsMu.RUnlock()

	// Find tag
	foundTag, exists := tags[req.Uuid]
	if !exists {
		return nil, status.Error(codes.NotFound, "tag not found")
	}

	log.Printf("Retrieved tag: %s (UUID: %s) by user: %s", foundTag.Name, foundTag.Uuid, req.Context.Username)

	return &ttt.GetTagResp{
		Tag: foundTag,
	}, nil
}

// ListTags lists all tags
func (s *TTTServer) ListTags(ctx context.Context, req *ttt.ListTagsReq) (*ttt.ListTagsResp, error) {
	log.Printf("ListTags called by user: %s", req.Context.Username)

	// Validate context
	if req.Context == nil || req.Context.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "context with username is required")
	}

	tagsMu.RLock()
	defer tagsMu.RUnlock()

	// Collect all tags
	var allTags []*tag.Tag
	for _, t := range tags {
		allTags = append(allTags, t)
	}

	log.Printf("Listed %d tags for user: %s", len(allTags), req.Context.Username)

	return &ttt.ListTagsResp{
		Tags: allTags,
	}, nil
}

// UpdateTag updates an existing tag
func (s *TTTServer) UpdateTag(ctx context.Context, req *ttt.UpdateTagReq) (*ttt.UpdateTagResp, error) {
	log.Printf("UpdateTag called by user: %s, UUID: %s, new name: %s", req.Context.Username, req.Uuid, req.Name)

	// Validate input
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "tag UUID cannot be empty")
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "tag name cannot be empty")
	}

	// Validate context
	if req.Context == nil || req.Context.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "context with username is required")
	}

	tagsMu.Lock()
	defer tagsMu.Unlock()

	// Find and update tag
	foundTag, exists := tags[req.Uuid]
	if !exists {
		return nil, status.Error(codes.NotFound, "tag not found")
	}

	oldName := foundTag.Name
	foundTag.Name = req.Name

	log.Printf("Updated tag: %s -> %s (UUID: %s) by user: %s", oldName, foundTag.Name, foundTag.Uuid, req.Context.Username)

	return &ttt.UpdateTagResp{
		Tag: foundTag,
	}, nil
}

// DeleteTag deletes a tag by UUID
func (s *TTTServer) DeleteTag(ctx context.Context, req *ttt.DeleteTagReq) (*ttt.DeleteTagResp, error) {
	log.Printf("DeleteTag called by user: %s, UUID: %s", req.Context.Username, req.Uuid)

	// Validate input
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "tag UUID cannot be empty")
	}

	// Validate context
	if req.Context == nil || req.Context.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "context with username is required")
	}

	tagsMu.Lock()
	defer tagsMu.Unlock()

	// Check if tag exists
	foundTag, exists := tags[req.Uuid]
	if !exists {
		return nil, status.Error(codes.NotFound, "tag not found")
	}

	// Delete tag
	delete(tags, req.Uuid)

	log.Printf("Deleted tag: %s (UUID: %s) by user: %s", foundTag.Name, foundTag.Uuid, req.Context.Username)

	return &ttt.DeleteTagResp{}, nil
}
