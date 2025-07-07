/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/moukhtar-youssef/Task_Tracker/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add <task>",
	Short:   "Add a new task to your task list",
	Long:    `Adds a new task to your task list. The task description should be provided as an argument.`,
	Example: "  Task_Tracker add \"Buy groceries\"",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internal.LpError(fmt.Errorf("Task description is required"))
			return
		}
		description := strings.Join(args[0:], " ")
		internal.AddTask(description)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
