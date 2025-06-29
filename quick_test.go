package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	contextpb "todo-time-tracker/proto/context"
	"todo-time-tracker/proto/ttt"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := ttt.NewTTTServiceClient(conn)
	ctx := context.Background()

	log.Println("🧪 Testing reflection-based interceptor...")

	// Test: Create a tag with valid context
	createReq := &ttt.CreateTagReq{
		Context: &contextpb.Context{
			Username: "reflection_test_user",
		},
		Name: "reflection_tag",
	}

	createResp, err := client.CreateTag(ctx, createReq)
	if err != nil {
		log.Printf("❌ CreateTag failed: %v", err)
	} else {
		log.Printf("✅ CreateTag succeeded: %s (UUID: %s)", createResp.Tag.Name, createResp.Tag.Uuid)
	}

	// Test: List tags
	listReq := &ttt.ListTagsReq{
		Context: &contextpb.Context{
			Username: "reflection_test_user",
		},
	}

	listResp, err := client.ListTags(ctx, listReq)
	if err != nil {
		log.Printf("❌ ListTags failed: %v", err)
	} else {
		log.Printf("✅ ListTags succeeded: found %d tags", len(listResp.Tags))
	}

	log.Println("🎉 Reflection-based interceptor test completed!")
}
