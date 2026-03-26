package alog

import (
	"fmt"
	"log"
	"sync"
)

// AsyncLogger provides non-blocking log output via a buffered channel.
// Messages are consumed by a background goroutine. When the channel is full,
// Log falls back to synchronous log.Printf (backpressure safety).
type AsyncLogger struct {
	ch    chan string
	wg    sync.WaitGroup
	once  sync.Once
	quiet bool
}

// New creates a new AsyncLogger with the given buffer size.
// If quiet is true, messages are consumed but not printed.
func New(bufferSize int, quiet bool) *AsyncLogger {
	l := &AsyncLogger{
		ch:    make(chan string, bufferSize),
		quiet: quiet,
	}
	l.wg.Add(1)
	go l.run()
	return l
}

// run is the consumer goroutine that drains the channel.
func (l *AsyncLogger) run() {
	defer l.wg.Done()
	for msg := range l.ch {
		if !l.quiet {
			log.Print(msg)
		}
	}
}

// Log formats and sends a message to the async channel.
// If the channel is full, falls back to synchronous log.Printf.
func (l *AsyncLogger) Log(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	select {
	case l.ch <- msg:
	default:
		log.Print(msg)
	}
}

// Close closes the channel and waits for all messages to be drained.
// Safe to call multiple times.
func (l *AsyncLogger) Close() {
	l.once.Do(func() {
		close(l.ch)
	})
	l.wg.Wait()
}
