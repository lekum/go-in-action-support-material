package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// interrupt channel reports a signal from the OS
	interrupt chan os.Signal
	// complete channel reports that processing is done
	complete chan error
	// timeout channel reports that the time has run out
	timeout <-chan time.Time
	// tasks holds a set of functions that are executed synchronously
	// in index order
	tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	// We want to receive all interrupt based signals
	signal.Notify(r.interrupt, os.Interrupt)

	// Run the different tasks on a different goroutine
	go func() {
		r.complete <- r.run()
	}()

	select {
	// Signaled when processing is done
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// gotInterrupt verifies if the interrupt signal has been issued
func (r *Runner) gotInterrupt() bool {
	select {
	// Signaled when an interrupt event is set
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	// Continue running as normal
	default:
		return false
	}
}
