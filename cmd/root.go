/*
Copyright © 2025 Moukhtar Youssef <moukhtar@example.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the base command
var RootCmd = &cobra.Command{
	Use:   "Task_Tracker",
	Short: "A CLI tool to manage your tasks",
	Long: `Task_Tracker is a simple and efficient command-line interface 
for tracking tasks. Add, edit, delete, and update your tasks easily.`,
}

// Execute runs the CLI
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle (not used)")
}
