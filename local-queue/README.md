# local-queue

Type-safe, in-memory local queue for development and testing.

## What it does

- **Type-safe**: `Queue[T]` so you push and consume a single message type.
- **Single worker**: One background goroutine per queue; messages processed one-by-one. No worker runs until a consumer is registered.
- **Retries**: If the consumer returns an error, the message is retried up to `MaxRetries` (default 3), then dropped.
- **No broker**: In-memory only; suitable for dev, tests, or single-process use.

## Installation

```bash
go get github.com/teathedev/pkg/local-queue
```

## Usage

```go
import "github.com/teathedev/pkg/localqueue"

type EmailJob struct {
	To      string
	Subject string
	Body    string
}

func main() {
	// Create queue with optional config
	q := localqueue.NewQueue[EmailJob]("emails", &localqueue.Options{
		MaxRetries: 3,
	})
	// Or default options: localqueue.NewQueue[EmailJob]("emails", nil)

	// Register consumer (starts the background worker)
	q.Consume(func(job EmailJob) error {
		err := sendEmail(job.To, job.Subject, job.Body)
		return err // non-nil => retry up to MaxRetries, then drop
	})

	// Push messages (non-blocking, thread-safe)
	q.Push(EmailJob{To: "user@example.com", Subject: "Hi", Body: "Hello"})
}
```

## API

| Type / Method                                       | Description                                                                                      |
| --------------------------------------------------- | ------------------------------------------------------------------------------------------------ |
| `NewQueue[T](name string, opts *Options) *Queue[T]` | Create queue. `opts` may be nil (MaxRetries=3).                                                  |
| `Options.MaxRetries`                                | Retries after consumer error (default 3).                                                        |
| `Queue[T].Push(msg T)`                              | Enqueue; non-blocking and safe for concurrent use.                                               |
| `Queue[T].Consume(consumer func(T) error)`          | Register consumer and start the single background worker. Must be called at most once per queue. |

## Behavior

- **No consumer**: No goroutine is started; `Push` only enqueues. No CPU use until `Consume` is called.
- **Order**: Messages are processed in FIFO order, one at a time.
- **Errors**: Consumer returns `error` → message is re-queued and retried until `MaxRetries`, then dropped (no logging inside the package).
