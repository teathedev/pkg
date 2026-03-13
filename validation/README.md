# validation

Struct validation using tags, with errors compatible with the errors package.

## What it does

- **ValidateStruct[T]**: Validates a struct using `go-playground/validator` and struct tags.
- **Errors**: On failure returns `*errors.CustomError` (400) with field-level details (`BadInputField`), using **json** tag names in error responses.
- **Tag name**: The validator uses the `json` tag for field names (e.g. `json:"email"` → "email" in the error).
- **RegisterValidation**: Lets you register custom validator functions for your own tags.

## Installation

```bash
go get github.com/teathedev/pkg/validation
```

Requires: `github.com/teathedev/pkg/errors`, `github.com/go-playground/validator/v10`.

## Usage

```go
import (
	"github.com/teathedev/pkg/errors"
	"github.com/teathedev/pkg/validation"
	validator "github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role"     validate:"oneof=admin user"`
}

func createUser(body *CreateUserRequest) *errors.CustomError {
	if err := validation.ValidateStruct(body); err != nil {
		return err // *errors.CustomError, status 400, with BadInputField list
	}
	// proceed with body
	return nil
}

// Register a custom validation tag
func init() {
	_ = validation.RegisterValidation("startsWithA", func(fl validator.FieldLevel) bool {
		s, ok := fl.Field().Interface().(string)
		if !ok {
			return false
		}
		return strings.HasPrefix(s, "A")
	})
}
```

## API

| Function                                                                                        | Description                                                                                                                                              |
| ----------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ValidateStruct[T any](data *T) *errors.CustomError`                                            | Validates `data`; nil on success. On failure returns 400-style CustomError with validation field details; field names come from `json` tags.             |
| `RegisterValidation(tag string, fn func(fl validator.FieldLevel) bool, callEvenIfNull ...bool)` | Registers a custom validation under `tag`. Delegates to `validator.RegisterValidation`; `callEvenIfNull` is passed through to control validation-on-nil. |

## Validator

Uses [go-playground/validator](https://github.com/go-playground/validator). Use standard tags such as `required`, `email`, `min`, `max`, `oneof`, etc.
