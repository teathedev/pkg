# logger

Structured, level-based logger with optional JSON output for production.

## What it does

- **Levels**: Trace, Info, Warning, Error, Fatal (Fatal calls `os.Exit(1)`).
- **Module** and **APP_NAME**: Each logger is bound to a module; app name comes from env `APP_NAME`.
- **Development**: Colored, human-readable lines (level, time, module, message, key-value params).
- **Production**: When `GO_ENV=production`, logs are JSON (app, module, level, message, date, and extra fields).
- **LogParams**: Extra data as `map[string]any` (or structs serialized as JSON).

## Installation

```bash
go get github.com/teathedev/pkg/logger
```

Requires: `github.com/teathedev/pkg/env`, `github.com/teathedev/pkg/utils`.

## Usage

```go
import "github.com/teathedev/pkg/logger"

func main() {
	log := logger.New("Payments")

	log.Info("charge started", logger.LogParams{"amount": 1000, "currency": "USD"})
	log.Warning("retry attempt", logger.LogParams{"attempt": 2})
	log.Error("charge failed", logger.LogParams{"error": err.Error(), "id": orderID})
	// log.Fatal("unrecoverable") // logs and os.Exit(1)

	// With struct (in production, serialized as JSON in "data")
	log.Info("order created", logger.LogParams{"data": myOrder})
}
```

Set environment for production JSON logs:

- `GO_ENV=production`
- `APP_NAME=my-service` (recommended; init warns if missing)

## API

| Function / Type             | Description                                                                                 |
| --------------------------- | ------------------------------------------------------------------------------------------- |
| `New(module string) Logger` | Creates a logger for that module (uses APP_NAME from env).                                  |
| `Logger`                    | Interface: `Trace`, `Info`, `Warning`, `Error`, `Fatal(message string, args ...LogParams)`. |
| `LogParams`                 | `map[string]any` for key-value or struct payloads.                                          |
