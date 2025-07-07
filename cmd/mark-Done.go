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

// markDoneCmd marks a task as done
var markDoneCmd = &cobra.Command{
	Use:     "mark-done <task-ID>",
	Short:   "Mark a task as done",
	Long:    `Marks the specified task as completed.`,
	Example: "  Task_Tracker mark-done 5",
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
		internal.MarkDone(id)
	},
}

func init() {
	RootCmd.AddCommand(markDoneCmd)
}
