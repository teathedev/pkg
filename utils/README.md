# utils

Small generic utilities used across teahub packages.

## What it does

- **IsStruct**: Reports whether a `reflect.Type` is a struct (e.g. for logging or validation).
- **MergeArrays**: Merges multiple slices of the same type into one.
- **GenerateRandomString**: Returns a random alphanumeric string of a given length.
- **GenerateRandomIntInRange**: Returns a random integer in [min, max] (inclusive).

No external dependencies beyond the standard library.

## Installation

```bash
go get github.com/teathedev/pkg/utils
```

## Usage

```go
import (
	"reflect"
	"github.com/teathedev/pkg/utils"
)

// Type check
func handle(v any) {
	t := reflect.TypeOf(v)
	if utils.IsStruct(t) {
		// v is a struct type
	}
}

// Merge slices
a, b, c := []int{1, 2}, []int{3}, []int{4, 5}
merged := utils.MergeArrays(a, b, c) // []int{1, 2, 3, 4, 5}

// Random string (A–Z, a–z, 0–9)
s := utils.GenerateRandomString(16) // e.g. "x7Kp2mNq9vLw4Rc1"

// Random int in range [min, max] inclusive
n := utils.GenerateRandomIntInRange(1, 100)
```

## API

| Function                                     | Description                                                    |
| -------------------------------------------- | -------------------------------------------------------------- |
| `IsStruct(t reflect.Type) bool`              | Returns true if `t.Kind() == reflect.Struct`.                  |
| `MergeArrays[T any](arrays ...[]T) []T`      | Concatenates all given slices into one.                        |
| `GenerateRandomString(length int) string`    | Random alphanumeric string (A–Z, a–z, 0–9) of length `length`. |
| `GenerateRandomIntInRange(min, max int) int` | Random int in [min, max] inclusive.                            |
