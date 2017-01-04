// Package pool manages a user defined set of resources
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool managea
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release places the new resource on the queue
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	// If the pool is closed, discard the resource
	if p.closed {
		r.Close()
		return
	}
	// Attempt to place the new resource on the queue
	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	// If the queue is already at capacity we close the resource
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close will shutdown the queue and close all existing resources
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	// Close the channel before we drain the channel of its resources
	close(p.resources)
	// Close the resources
	for r := range p.resources {
		r.Close()
	}
}
