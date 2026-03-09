# errors

Structured, HTTP-aware errors for APIs and applications.

## What it does

- **CustomError**: JSON-serializable error with module, message, optional params, and HTTP status.
- Helpers for common HTTP errors (400, 401, 403, 404, 415, 500).
- **GetStatus()** for framework integration (e.g. Huma `StatusError`).
- **IsNotFound(err)** to detect 404-style custom errors.
- **APIErrorResponse** / **APIErrorSimple** for OpenAPI/400 vs 401/403/404/500 response shapes.

## Installation

```bash
go get github.com/teathedev/pkg/errors
```

## Usage

```go
import "github.com/teathedev/pkg/errors"

// Generic error (500)
err := errors.New("Payments", "charge failed")

// With HTTP status (e.g. for Huma)
err := errors.NewWithStatus("API", "Invalid ID", 400)

// Standard HTTP errors
err := errors.NewUnauthorizedError("Auth")
err := errors.NewAccessDeniedError("Admin")
err := errors.NewNotFoundError("Users", "user not found")
err := errors.NewUnsupportedFileFormatError("Upload")

// Validation / bad input (400) with field details
err := errors.NewBadInput("Validation", []errors.BadInputField{
	{Field: "email", Condition: errors.BadInputConditionNotValid, Value: "bad"},
})

// Check for 404
if errors.IsNotFound(err) {
	// handle not found
}

// Use in API responses (Huma etc.)
if customErr, ok := err.(*errors.CustomError); ok {
	status := customErr.GetStatus()
	// return status and customErr (serializes as JSON)
}
```

## API overview

| Function / Type                                                   | Description                                   |
| ----------------------------------------------------------------- | --------------------------------------------- |
| `New(module, message string) *CustomError`                        | 500 error.                                    |
| `NewWithStatus(module, message string, status int) *CustomError`  | Custom status.                                |
| `NewBadInput(module string, params []BadInputField) *CustomError` | 400 validation error.                         |
| `NewNotFoundError(module, message string) *CustomError`           | 404.                                          |
| `NewUnauthorizedError(module string) *CustomError`                | 401.                                          |
| `NewAccessDeniedError(module string) *CustomError`                | 403.                                          |
| `NewUnsupportedFileFormatError(module string) *CustomError`       | 415.                                          |
| `IsNotFound(err error) bool`                                      | True if err is \*CustomError with status 404. |
| `CustomError.GetStatus() int`                                     | HTTP status for the error.                    |
| `APIErrorResponse`                                                | 400 body shape (message + error array).       |
| `APIErrorSimple`                                                  | 401/403/404/500 body shape (message only).    |
