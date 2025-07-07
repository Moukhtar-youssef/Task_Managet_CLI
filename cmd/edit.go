/*
Copyright © 2025 Moukhtar Youssef <moukhtar@example.com>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/moukhtar-youssef/Task_Tracker/internal"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit <task-ID> <new-description>",
	Short:   "Edit the description of an existing task",
	Long:    `Modifies the description of an existing task by its ID.`,
	Example: "  Task_Tracker edit 3 \"Read a new book\"",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			internal.LpError(fmt.Errorf("Usage: edit <task-ID> <new-description>"))
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			internal.LpError(fmt.Errorf("Invalid task ID: %v", err))
			return
		}
		newDesc := strings.Join(args[1:], " ")
		internal.Edit(id, newDesc)
	},
}

func init() {
	RootCmd.AddCommand(editCmd)
}
