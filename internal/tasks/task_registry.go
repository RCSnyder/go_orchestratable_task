package tasks

var taskRegistry = map[string]OrchestratableTask{}

// RegisterTask registers a task with a given name.
func RegisterTask(name string, task OrchestratableTask) {
    taskRegistry[name] = task
}

// GetTask retrieves a registered task by name.
func GetTask(name string) (OrchestratableTask, bool) {
    task, exists := taskRegistry[name]
    return task, exists
}

// Initialize and register all tasks
func init() {
    RegisterTask("ExampleTask", NewExampleTask())
}
