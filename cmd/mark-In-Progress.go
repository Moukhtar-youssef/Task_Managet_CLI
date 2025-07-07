/*
Copyright © 2025 Moukhtar Youssef <moukhtar@example.com>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/moukhtar-youssef/Task_Tracker/internal"
	"github.com/spf13/cobra"
)

// markInProgressCmd marks a task as in-progress
var markInProgressCmd = &cobra.Command{
	Use:     "mark-in-progress <task-ID>",
	Short:   "Mark a task as in progress",
	Long:    `Updates the status of a task to 'in-progress'.`,
	Example: "  Task_Tracker mark-in-progress 4",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internal.LpError(fmt.Errorf("Task ID is required"))
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			internal.LpError(fmt.Errorf("Invalid task ID: %v", err))
			return
		}
		internal.MarkInprogress(id)
	},
}

func init() {
	RootCmd.AddCommand(markInProgressCmd)
}
