package rpc

import (
	"context"
	"fmt"

	contextpb "todo-time-tracker/proto/go/context"
	"todo-time-tracker/proto/go/ttt"
)

// CreateTag creates a new tag via gRPC
func (c *TTTClient) CreateTag(ctx context.Context, name string) (*ttt.CreateTagResp, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &ttt.CreateTagReq{
		Context: &contextpb.Context{
			UserName: c.username,
		},
		Name: name,
	}

	resp, err := c.client.CreateTag(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create tag: %v", err)
	}

	return resp, nil
}

// GetTag retrieves a tag by UUID via gRPC
func (c *TTTClient) GetTag(ctx context.Context, uuid string) (*ttt.GetTagResp, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &ttt.GetTagReq{
		Context: &contextpb.Context{
			UserName: c.username,
		},
		Uuid: uuid,
	}

	resp, err := c.client.GetTag(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get tag: %v", err)
	}

	return resp, nil
}
