# teahub/pkg

Go packages for use across TeaHub projects. Each package can be installed and used independently.

**Module:** `github.com/teathedev/pkg`

## Packages

| Package                      | Description                                                              |
| ---------------------------- | ------------------------------------------------------------------------ |
| [env](./env)                 | Environment variable helpers (string, number) with defaults.             |
| [errors](./errors)           | Structured, HTTP-aware errors (CustomError, 400/401/403/404/415/500).    |
| [jwt](./jwt)                 | JWT/JWK: key generation, encrypt at rest, PEM, JWK for OIDC discovery.   |
| [logger](./logger)           | Level-based logger (Trace/Info/Warning/Error/Fatal), JSON in production. |
| [local-queue](./local-queue) | Type-safe in-memory queue for dev; one worker, retries, no broker.       |
| [utils](./utils)             | Small helpers (e.g. IsStruct).                                           |
| [validation](./validation)   | Struct validation with go-playground/validator → CustomError (400).      |

## Installation

Use the whole module (e.g. in your application):

```bash
go get github.com/teathedev/pkg
```

Or add a specific package; the module is still required:

```bash
go get github.com/teathedev/pkg/errors
go get github.com/teathedev/pkg/env
go get github.com/teathedev/pkg/jwt
go get github.com/teathedev/pkg/logger
go get github.com/teathedev/pkg/local-queue
go get github.com/teathedev/pkg/utils
go get github.com/teathedev/pkg/validation
```

## Usage

Import by package path:

```go
import (
	"github.com/teathedev/pkg/env"
	"github.com/teathedev/pkg/errors"
	"github.com/teathedev/pkg/logger"
)
```

See each package’s README for what it does and how to use it.

## Requirements

- Go 1.25.6 or later (see [go.mod](./go.mod)).

## License

See [LICENSE](./LICENSE).
