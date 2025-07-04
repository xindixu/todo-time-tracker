package cli

import (
	"fmt"
	"log"
	"os"
	"todo-time-tracker/client/rpc"

	"github.com/spf13/cobra"
)

var cliRPCClient *rpc.TTTClient

// CreateRPCClient creates a new gRPC client for the server
func CreateRPCClient() {
	var err error
	cliRPCClient, err = rpc.NewTTTClient("", "") // Use defaults
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to server: %v\n", err)
		os.Exit(1)
	}
}

// CloseRPCClient closes the gRPC client
func CloseRPCClient() {
	err := cliRPCClient.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error closing gRPC client: %v\n", err)
		os.Exit(1)
	}
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ttt",
	Short: "Simple CLI tool to help you keep track of the tasks to do and the time spent on a task",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	log.Println("Executing root command")
	err := RootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing root command: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo-time-tracker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
