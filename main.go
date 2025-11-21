package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sudosantos27/go-task-tracker/task"
)

func main() {
	// Check if any command was passed
	if len(os.Args) < 2 {
		fmt.Println("Usage: task <command> [arguments]")
		fmt.Println("Available commands: add, list, complete, delete, info")
		os.Exit(1)
	}

	// The first argument is the subcommand (add, list, etc.)
	switch os.Args[1] {
	case "add":
		// Define specific flags for the 'add' command
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("title", "", "Task title")
		description := addCmd.String("desc", "", "Task description")

		// Parse arguments starting from the second one (os.Args[2:])
		addCmd.Parse(os.Args[2:])

		if *title == "" {
			fmt.Println("Error: Title is required. Usage: task add -title=\"My task\"")
			os.Exit(1)
		}

		t, err := task.Add(*title, *description)
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task added: %d: %s\n", t.ID, t.Title)

	case "list":
		tasks, err := task.List()
		if err != nil {
			fmt.Printf("Error loading tasks: %v\n", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("No pending tasks.")
			return
		}

		fmt.Println("Task list:")
		for _, t := range tasks {
			status := "[ ]"
			if t.Done {
				status = "[x]"
			}
			// Format: [ ] 1: Buy coffee (created: 2025-11-21 13:00)
			fmt.Printf("%s %d: %s (created: %s)\n", status, t.ID, t.Title, t.CreatedAt.Format("2006-01-02 15:04"))
		}

	case "complete":
		completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
		id := completeCmd.Int("id", 0, "ID of the task to complete")

		completeCmd.Parse(os.Args[2:])

		if *id == 0 {
			fmt.Println("Error: ID is required. Usage: task complete -id=1")
			os.Exit(1)
		}

		if err := task.Complete(*id); err != nil {
			fmt.Printf("Error completing task: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task %d marked as completed.\n", *id)

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "ID of the task to delete")

		deleteCmd.Parse(os.Args[2:])

		if *id == 0 {
			fmt.Println("Error: ID is required. Usage: task delete -id=1")
			os.Exit(1)
		}

		if err := task.Delete(*id); err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task %d deleted.\n", *id)

	case "info":
		infoCmd := flag.NewFlagSet("info", flag.ExitOnError)
		id := infoCmd.Int("id", 0, "ID of the task to view")

		infoCmd.Parse(os.Args[2:])

		if *id == 0 {
			fmt.Println("Error: ID is required. Usage: task info -id=1")
			os.Exit(1)
		}

		handleInfo(*id)

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func handleInfo(id int) {
	tasks, err := task.List()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	for _, t := range tasks {
		if t.ID == id {
			fmt.Printf("ID: %d\n", t.ID)
			fmt.Printf("Title: %s\n", t.Title)
			fmt.Printf("Description: %s\n", t.Description)
			fmt.Printf("Status: %v\n", t.Done)
			fmt.Printf("Created At: %s\n", t.CreatedAt.Format("2006-01-02 15:04:05"))
			if t.CompletedAt != nil {
				fmt.Printf("Completed At: %s\n", t.CompletedAt.Format("2006-01-02 15:04:05"))
			}
			return
		}
	}

	fmt.Printf("Error: Task with ID %d not found\n", id)
	os.Exit(1)
}
