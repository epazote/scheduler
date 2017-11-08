package scheduler

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// scheduler private struct
type scheduler struct {
	t    <-chan time.Time
	quit chan struct{}
	f    func()
}

// Scheduler map of schedulers
type Scheduler struct {
	sync.Map
}

// New returns a new scheduler
func New() *Scheduler {
	return &Scheduler{}
}

// AddScheduler calls a function every X seconds.
func (s *Scheduler) AddScheduler(name string, interval int, f func()) {
	e := time.Duration(interval) * time.Second

	// create a new scheduler
	task := scheduler{
		t:    time.NewTicker(e).C,
		quit: make(chan struct{}),
		f:    f,
	}

	// stop scheduler if exist or add a new one
	if sk, ok := s.LoadOrStore(name, task); ok {
		close(sk.(scheduler).quit)
	}

	// Create the scheduler in a goroutine running forever until it quits
	go func(s scheduler) {
		for {
			select {
			case <-s.t:
				s.f()
			case <-s.quit:
				return
			}
		}
	}(task)
}

// Stop ends a specified scheduler.
func (s *Scheduler) Stop(name string) error {
	sk, ok := s.Load(name)
	if !ok {
		return fmt.Errorf("Scheduler: %s, does not exist.", name)
	}
	close(sk.(scheduler).quit)
	return nil
}

// StopAll ends all schedulers.
func (s *Scheduler) StopAll() {
	close := func(key, value interface{}) bool {
		close(value.(scheduler).quit)
		log.Printf("Stoping: %s", key)
		return true
	}
	s.Range(close)
}
