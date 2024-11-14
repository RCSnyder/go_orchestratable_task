```
my_project/
├── cmd/
│   └── ot/
│       ├── main.go             # Main entry point for the CLI tool
│       └── run.go              # CLI run command implementation
├── internal/
│   └── tasks/
│       ├── base_task.go        # Base task interface and struct
│       ├── task_registry.go    # Task registry for auto-discovery
│       ├── example_task.go     # Example task definition
├── pkg/
│   ├── utils/
│   │   ├── s3.go               # Utility for fetching data from S3
│   │   └── json_parser.go      # Utility for parsing JSON strings
├── go.mod                      # Go module file for dependency management
└── README.md                   # Instructions and documentation
```