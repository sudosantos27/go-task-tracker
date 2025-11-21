# Go Task Tracker

A simple, robust Command Line Interface (CLI) task manager built with Go. This project demonstrates core Go concepts including structs, slices, JSON persistence, file I/O, and package organization.

## Features

- **Task Management**: Add, list, complete, and delete tasks.
- **Rich Details**: Support for task titles, descriptions, and completion timestamps.
- **Persistence**: Automatically saves data to a local `tasks.json` file.
- **Filtering**: Filter tasks by status (todo/done).
- **Modular Design**: Clean separation of concerns using a dedicated `task` package.

## Prerequisites

- [Go](https://go.dev/dl/) installed (version 1.23 or higher recommended).

## Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/sudosantos27/go-task-tracker.git
    cd go-task-tracker
    ```

2.  (Optional) Build the binary:
    ```bash
    go build -o task cmd/task-cli/main.go
    ```
    *If you build the binary, replace `go run cmd/task-cli/main.go` with `./task` in the examples below.*

## Usage

The general syntax is:
```bash
go run cmd/task-cli/main.go <command> [flags]
```

### 1. Add a Task
Create a new task with a title and an optional description.

```bash
go run cmd/task-cli/main.go add -title="Learn Go" -desc="Study structs and interfaces"
```

### 2. List Tasks
View all tasks.

```bash
go run cmd/task-cli/main.go list
```

**Filter by status:**
```bash
go run cmd/task-cli/main.go list -status=todo  # Show only pending tasks
go run cmd/task-cli/main.go list -status=done  # Show only completed tasks
```

### 3. View Task Details
Get detailed information about a specific task, including its description and timestamps.

```bash
go run cmd/task-cli/main.go info -id=1
```

### 4. Complete a Task
Mark a task as done. This records the completion time.

```bash
go run cmd/task-cli/main.go complete -id=1
```

### 5. Delete a Task
Permanently remove a task.

```bash
go run cmd/task-cli/main.go delete -id=1
```

## Project Structure

```
go-task-tracker/
├── cmd/
│   ├── task-cli/   # CLI Entry point.
│   │   └── main.go
│   └── task-api/   # API Entry point.
│       └── main.go
├── go.mod          # Go module definition.
├── tasks.json      # Local storage for tasks (created automatically).
├── internal/       # Internal application code.
│   ├── task/       # Domain model and repository interface.
│   ├── store/      # Persistence implementations.
│   └── api/        # HTTP Handlers and middleware.
└── task/           # (Legacy) Business logic package for CLI.
    ├── task.go
    └── task_test.go
```

## API Usage

The project includes a REST API server.

### 1. Start the Server
```bash
go run cmd/task-api/main.go
```
The server will start on `http://localhost:8080`.

To stop the server, press `Ctrl+C` in your terminal.

### 2. Check Health
```bash
curl http://localhost:8080/health
# Output: {"status":"ok"}
```

### 3. Create a Task
```bash
curl -X POST -H "Content-Type: application/json" \
     -d '{"title":"Learn API Testing", "description":"Use httptest package"}' \
     http://localhost:8080/tasks
```

### 4. List Tasks
```bash
curl http://localhost:8080/tasks
```

## Running Tests

### CLI Tests
```bash
go test ./task -v
```

### API Tests
```bash
go test ./internal/api -v
```
