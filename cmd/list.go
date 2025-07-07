/*
Copyright © 2025 Moukhtar Youssef <moukhtar@example.com>
*/
package cmd

import (
	"fmt"

	"github.com/moukhtar-youssef/Task_Tracker/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list [status]",
	Short:   "List all tasks or filter by status",
	Long:    `Lists all tasks in your task list. Optionally filter by status: todo, in-progress, or done.`,
	Example: "  Task_Tracker list\n  Task_Tracker list done",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internal.ListTasks()
		} else if len(args) > 1 {
			internal.LpError(fmt.Errorf("Only one status can be provided: 'todo', 'in-progress', or 'done'"))
			return
		} else {
			internal.ListFilter(args[0])
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
