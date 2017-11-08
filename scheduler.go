package scheduler

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// scheduler private struct
type scheduler struct {
	f    func()
	quit chan struct{}
	t    *time.Ticker
}

// Scheduler map of schedulers
type Scheduler struct {
	sync.Map
}

// New returns a new scheduler
func New() *Scheduler {
	return &Scheduler{}
}

// AddScheduler calls a function every X defined interval
func (s *Scheduler) AddScheduler(name string, interval time.Duration, f func()) {
	// create a new scheduler
	task := scheduler{
		f:    f,
		quit: make(chan struct{}),
		t:    time.NewTicker(interval),
	}

	// stop scheduler if exist
	if sk, ok := s.Load(name); ok {
		close(sk.(scheduler).quit)
	}

	// add a new task
	s.Store(name, task)

	// create the scheduler in a goroutine running forever until it quits
	go func() {
		for {
			select {
			case <-task.t.C:
				task.f()
			case <-task.quit:
				task.t.Stop()
				return
			}
		}
	}()
}

// Stop ends a specified scheduler.
func (s *Scheduler) Stop(name string) error {
	sk, ok := s.Load(name)
	if !ok {
		return fmt.Errorf("Scheduler: %s, does not exist.", name)
	}
	close(sk.(scheduler).quit)
	s.Delete(name)
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
