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
	Use:   "add <task>",
	Short: "Add a task to the todo list",
	Long:  `Add a task to the todo list it is done like that 'add <task>'`,
	Example: `
	Todo add Todo1 
	Todo add Buy milk
	`,
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
	rootCmd.AddCommand(addCmd)
}
