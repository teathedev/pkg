# env

Environment variable helpers with safe defaults.

## What it does

- Reads string and integer environment variables.
- Returns a default value when the variable is unset or invalid (e.g. non-numeric for `GetNumber`).
- No external dependencies beyond the standard library.

## Installation

```bash
go get github.com/teathedev/pkg/env
```

## Usage

```go
import "github.com/teathedev/pkg/env"

func main() {
	// String with default
	port := env.GetString("PORT", "8080")
	dbURL := env.GetString("DATABASE_URL", "")

	// Integer with default (invalid or empty => defaultValue)
	workers := env.GetNumber("WORKERS", 4)
	timeout := env.GetNumber("TIMEOUT_SECONDS", 30)
}
```

## API

| Function                                      | Description                                                    |
| --------------------------------------------- | -------------------------------------------------------------- |
| `GetString(key, defaultValue string) string`  | Returns `os.Getenv(key)` or `defaultValue` if empty.           |
| `GetNumber(key string, defaultValue int) int` | Parses key as int; returns `defaultValue` if unset or invalid. |
