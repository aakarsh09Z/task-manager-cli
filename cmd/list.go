package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of tasks",
	Long:  `This command gives the list of all the current tasks which were added and not removed.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := readTasksFromFile()
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		fmt.Println()
		for _, task := range tasks {
			fmt.Printf("Title: %s \nDescription: %s \nTimestamp: %s \n\n",
				task.Title, task.Description, task.Timestamp.Format("2006-01-02 15:04:05"))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
