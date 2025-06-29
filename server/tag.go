package server

import (
	"context"
	"log"
	"sync"

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

// CreateTag creates a new tag
func (s *TTTServer) CreateTag(ctx context.Context, req *ttt.CreateTagReq) (*ttt.CreateTagResp, error) {
	username := getUsername(ctx)

	// Validate input (authentication is handled by interceptor)
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "tag name cannot be empty")
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

	log.Printf("CreateTag: Successfully created tag: %s (UUID: %s) by user: %s", newTag.Name, newTag.Uuid, username)

	return &ttt.CreateTagResp{
		Tag: newTag,
	}, nil
}

// GetTag retrieves a tag by UUID
func (s *TTTServer) GetTag(ctx context.Context, req *ttt.GetTagReq) (*ttt.GetTagResp, error) {
	username := getUsername(ctx)

	// Validate input
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "tag UUID cannot be empty")
	}

	tagsMu.RLock()
	defer tagsMu.RUnlock()

	// Find tag
	foundTag, exists := tags[req.Uuid]
	if !exists {
		return nil, status.Error(codes.NotFound, "tag not found")
	}

	log.Printf("GetTag: Successfully retrieved tag: %s (UUID: %s) for user: %s", foundTag.Name, foundTag.Uuid, username)

	return &ttt.GetTagResp{
		Tag: foundTag,
	}, nil
}

// ListTags lists all tags
func (s *TTTServer) ListTags(ctx context.Context, req *ttt.ListTagsReq) (*ttt.ListTagsResp, error) {
	username := getUsername(ctx)

	tagsMu.RLock()
	defer tagsMu.RUnlock()

	// Collect all tags
	var allTags []*tag.Tag
	for _, t := range tags {
		allTags = append(allTags, t)
	}

	log.Printf("ListTags: Successfully listed %d tags for user: %s", len(allTags), username)

	return &ttt.ListTagsResp{
		Tags: allTags,
	}, nil
}

// UpdateTag updates an existing tag
func (s *TTTServer) UpdateTag(ctx context.Context, req *ttt.UpdateTagReq) (*ttt.UpdateTagResp, error) {
	username := getUsername(ctx)

	// Validate input
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "tag UUID cannot be empty")
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "tag name cannot be empty")
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

	log.Printf("UpdateTag: Successfully updated tag: %s -> %s (UUID: %s) by user: %s", oldName, foundTag.Name, foundTag.Uuid, username)

	return &ttt.UpdateTagResp{
		Tag: foundTag,
	}, nil
}

// DeleteTag deletes a tag by UUID
func (s *TTTServer) DeleteTag(ctx context.Context, req *ttt.DeleteTagReq) (*ttt.DeleteTagResp, error) {
	username := getUsername(ctx)

	// Validate input
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "tag UUID cannot be empty")
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

	log.Printf("DeleteTag: Successfully deleted tag: %s (UUID: %s) by user: %s", foundTag.Name, foundTag.Uuid, username)

	return &ttt.DeleteTagResp{}, nil
}
