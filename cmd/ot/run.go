package main

import (
    "encoding/json"
    "fmt"
    "log"
    "go_orchestratable_task/internal/tasks"
    "github.com/spf13/cobra"
)

var taskName string
var kwargs string

func newRunCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "run",
        Short: "Run a specific task",
        Run:   runTask,
    }

    cmd.Flags().StringVar(&taskName, "task", "", "Name of the task to run")
    cmd.Flags().StringVar(&kwargs, "kwargs", "", "JSON string with task arguments")
    cmd.MarkFlagRequired("task")
    cmd.MarkFlagRequired("kwargs")

    return cmd
}

func runTask(cmd *cobra.Command, args []string) {
    // Retrieve the task from the registry
    task, exists := tasks.GetTask(taskName)
    if !exists {
        log.Fatalf("Task %s not found", taskName)
    }

    // Process kwargs as per the specific task's requirements
    var input tasks.OrchestratableTaskInputArgs
    switch taskName {
    case "ExampleTask":
        var exampleInput tasks.ExampleInput
        if err := json.Unmarshal([]byte(kwargs), &exampleInput); err != nil {
            log.Fatalf("Failed to parse kwargs: %v", err)
        }
        input = exampleInput
    default:
        log.Fatalf("No input type found for task %s", taskName)
    }

    // Execute the task with the parsed input
    output, err := task.Run(input)
    if err != nil {
        log.Fatalf("Error executing task %s: %v", taskName, err)
    }

    fmt.Printf("Output: %+v\n", output)
}

