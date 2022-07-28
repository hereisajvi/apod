package worker

import (
	"context"
	"log"
	"time"
)

const (
	duration = 12 * time.Hour
	layout   = "2006-01-02"
)

type Option interface {
	apply(worker *Worker)
}

type OptionFn func(worker *Worker)

func (fn OptionFn) apply(worker *Worker) {
	fn(worker)
}

func WithDuration(duration time.Duration) OptionFn {
	return func(worker *Worker) {
		worker.ticker = time.NewTicker(duration)
	}
}

// Worker contains all methods for receiving pictures from the APOD API by specific period.
type Worker struct {
	service PictureService
	ticker  *time.Ticker
	done    chan struct{}
}

// New is a constructor for Worker.
func New(service PictureService, opts ...Option) *Worker {
	worker := &Worker{
		service: service,
		ticker:  time.NewTicker(duration),
	}

	for _, opt := range opts {
		opt.apply(worker)
	}

	return worker
}

// Run starts worker and ticker and retrieves picture from the APOD API on each tick.
func (w Worker) Run(ctx context.Context) {
	now := time.Now()

	fn := func() {
		_, err := w.service.GetByDate(ctx, now.Format(layout))
		if err != nil {
			log.Printf("could not get picture by date: %v\n", err)
		}
	}

	fn()

	for {
		select {
		case <-w.ticker.C:
			fn()
		case <-w.done:
			log.Println("Stopping shudow worker...")
			return
		}
	}
}

// Stop stops worker.
func (w Worker) Stop() {
	w.done <- struct{}{}
}
