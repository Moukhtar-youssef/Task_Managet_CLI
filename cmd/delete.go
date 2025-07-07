/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/moukhtar-youssef/Task_Tracker/internal"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete <task-ID>",
	Short:   "Delete a task by its ID",
	Long:    `Deletes a task from your list by specifying its numeric ID.`,
	Example: "  Task_Tracker delete 2",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internal.LpError(fmt.Errorf("Can't be empty"))
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			internal.LpError(fmt.Errorf("Invalid task ID: %v", err))
		}
		internal.DeleteTask(id)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
