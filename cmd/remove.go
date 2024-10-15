package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task by title or index",
	Long:  `This command removes a task from the task manager using its title or index`,
	Run: func(cmd *cobra.Command, args []string) {
		allTasks, _ := cmd.Flags().GetBool("all")
		tasks := readTasksFromFile()
		if allTasks {
			if len(tasks) == 0 {
				fmt.Println("No tasks to remove.")
				return
			}

			err := os.Remove(filePath)
			if err != nil {
				fmt.Println("Error removing tasks:", err)
				return
			}

			fmt.Println("Removed")
			return
		}
		if len(args) < 1 {
			fmt.Println("Insufficient number of arguments")
			return
		}

		if len(tasks) == 0 {
			fmt.Println("Empty tasks file")
			return
		}

		arg := args[0]
		if index, err := strconv.Atoi(arg); err == nil && index >= 0 && index < len(tasks) {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Printf("Removed\n")
		} else {
			removed := false
			for i, task := range tasks {
				if strings.EqualFold(task.Title, arg) {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Printf("Removed\n")
					removed = true
					break
				}
			}
			if !removed {
				fmt.Printf("Title '%s' not found.\n", arg)
				return
			}
		}

		writeTasksToFile(tasks)
	},
}

func init() {
	removeCmd.Flags().BoolP("all", "a", false, "Remove all tasks")
	rootCmd.AddCommand(removeCmd)
}
