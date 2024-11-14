package tasks

import "fmt"

// ExampleInput represents the expected input for ExampleTask
type ExampleInput struct {
    Field string `json:"Field"`
}

// Validate ensures that ExampleInput meets requirements
func (i ExampleInput) Validate() error {
    if i.Field == "" {
        return fmt.Errorf("Field cannot be empty")
    }
    return nil
}

// ExampleOutput represents the output of ExampleTask
type ExampleOutput struct {
    Result string
}

// ExampleTask implements OrchestratableTask
type ExampleTask struct {
    BaseOrchestratableTask
}

func NewExampleTask() *ExampleTask {
    return &ExampleTask{
        BaseOrchestratableTask: NewBaseOrchestratableTask(ExampleInput{}, ExampleOutput{}),
    }
}

func (t *ExampleTask) Run(input OrchestratableTaskInputArgs) (OrchestratableTaskOutput, error) {
    // Assert input to ExampleInput type
    exampleInput, ok := input.(ExampleInput)
    if !ok {
        return nil, fmt.Errorf("invalid input type")
    }

    // Validate input
    if err := exampleInput.Validate(); err != nil {
        return nil, err
    }

    // Perform the task
    output := ExampleOutput{
        Result: "Processed: " + exampleInput.Field,
    }
    return output, nil
}
