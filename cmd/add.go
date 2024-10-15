package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	title       string
	description string
	filePath    = "data/tasks.json"
)

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the list",
	Long:  `This command adds a task to the list with description. To add a description you can use the flag -d`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		if title == "" {
			fmt.Print("Title: ")
			inputTitle, _ := reader.ReadString('\n')
			title = strings.TrimSpace(inputTitle)
		}

		if description == "" {
			fmt.Print("Description: ")
			desc, _ := reader.ReadString('\n')
			description = strings.TrimSpace(desc)
		}

		addTask(title, description)
		fmt.Printf("Task added: %s - %s\n", title, description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the task")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the task")
}

func addTask(title, description string) {
	tasks := readTasksFromFile()

	newTask := Task{
		Title:       title,
		Description: description,
		Timestamp:   time.Now(),
	}

	tasks = append(tasks, newTask)

	writeTasksToFile(tasks)
}

func readTasksFromFile() []Task {
	var tasks []Task

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return tasks
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading tasks file:", err)
		return tasks
	}

	err = json.Unmarshal(fileContent, &tasks)
	if err != nil {
		fmt.Println("Error parsing tasks:", err)
	}

	return tasks
}

func writeTasksToFile(tasks []Task) {
	fileContent, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling tasks:", err)
		return
	}

	os.MkdirAll("data", os.ModePerm)

	err = os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file:", err)
	}
}
