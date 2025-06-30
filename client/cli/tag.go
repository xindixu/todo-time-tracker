package cli

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tags",
	Long:  "Commands for managing tags in the time tracking system.",
}

var tagAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new tag",
	Long:  "Add a new tag to the system. You will be prompted to enter the tag name.",
	Run: func(cmd *cobra.Command, args []string) {
		addTag()
	},
}

func init() {
	RootCmd.AddCommand(tagCmd)
	tagCmd.AddCommand(tagAddCmd)
}

func addTag() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter tag name: ")
	tagName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return
	}

	// Trim whitespace and newline
	tagName = strings.TrimSpace(tagName)

	if tagName == "" {
		fmt.Println("Tag name cannot be empty")
		return
	}

	log.Println("Tag will be created with name:", tagName)

	// Create the tag via gRPC
	CreateRPCClient()

	resp, err := cliRPCClient.CreateTag(context.Background(), tagName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating tag: %v\n", err)
		return
	}
	CloseRPCClient()

	// Display success message
	fmt.Printf("✅ Successfully created tag: '%s' (UUID: %s)\n", resp.Tag.Name, resp.Tag.Uuid)
}
