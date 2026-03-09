// Package localqueue provides a type-safe local in-memory queue for development.
// A single background worker processes messages one-by-one. No background task
// runs until a consumer is registered.
package localqueue

import (
	"container/list"
	"sync"
)

// Options configures queue behaviour.
type Options struct {
	// MaxRetries is the number of retries after a consumer returns an error (default 3).
	MaxRetries int
}

func (o *Options) maxRetries() int {
	if o != nil && o.MaxRetries > 0 {
		return o.MaxRetries
	}
	return 3
}

type item[T any] struct {
	payload T
	retries int
}

// Queue is a type-safe queue that processes messages one-by-one in a single background goroutine.
// The background task is started only when Consume is called.
type Queue[T any] struct {
	name    string
	opts    *Options
	mu      sync.Mutex
	list    *list.List
	cond    *sync.Cond
	consume func(T) error
	running bool
}

// NewQueue creates a new queue with the given name and options.
// Options may be nil to use defaults (MaxRetries: 3).
func NewQueue[T any](name string, opts *Options) *Queue[T] {
	q := &Queue[T]{
		name: name,
		opts: opts,
		list: list.New(),
	}
	q.cond = sync.NewCond(&q.mu)
	return q
}

// Push enqueues a message. It is non-blocking and thread-safe.
func (q *Queue[T]) Push(msg T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.list.PushBack(&item[T]{payload: msg, retries: 0})
	q.cond.Signal()
}

// Consume registers the consumer and starts the single background worker.
// It must be called at most once per queue. Messages are processed one-by-one.
// If the consumer returns an error, the message is retried up to MaxRetries, then dropped.
func (q *Queue[T]) Consume(consumer func(T) error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.consume != nil {
		panic("localqueue: Consume already called")
	}
	q.consume = consumer
	if !q.running {
		q.running = true
		go q.run()
	}
}

func (q *Queue[T]) run() {
	maxRetries := q.opts.maxRetries()

	for {
		q.mu.Lock()
		for q.list.Len() == 0 {
			q.cond.Wait()
		}
		front := q.list.Front()
		q.list.Remove(front)
		it := front.Value.(*item[T])
		q.mu.Unlock()

		err := q.consume(it.payload)
		if err != nil {
			q.mu.Lock()
			if it.retries < maxRetries {
				it.retries++
				q.list.PushBack(it)
				q.cond.Signal()
			}
			q.mu.Unlock()
		}
	}
}
