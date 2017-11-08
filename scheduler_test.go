package scheduler

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

type count struct {
	i int64
}

func (c *count) Add(v int64) {
	atomic.AddInt64(&c.i, v)
}
func (c *count) Del(v int64) {
	atomic.AddInt64(&c.i, -v)
}
func (c *count) Get() int64 {
	return atomic.LoadInt64(&c.i)
}

func TestScheduler(t *testing.T) {
	sk := New()
	c := &count{}
	sk.AddScheduler("print", time.Second, func() { fmt.Println(c.Get()); c.Add(1) })
	select {
	case <-time.After(3 * time.Second):
		sk.StopAll()
	}
	if c.Get() <= 1 {
		t.Error("Expecting c > 0")
	}
	every := time.Millisecond
	sk.AddScheduler("print", every, func() { fmt.Println(c.Get()); c.Add(1) })
	time.Sleep(2 * time.Millisecond)
	if c.Get() < 1 {
		t.Fatalf("Expecting c > 0, got: %v", c.Get())
	}
	sk.StopAll()
}

func TestStopError(t *testing.T) {
	sk := New()
	c := &count{}
	sk.AddScheduler("print", 1, func() { c.Add(1) })
	err := sk.Stop("none")
	if err == nil {
		t.Error("Expecting error")
	}
}

func TestStop(t *testing.T) {
	sk := New()
	c := &count{}
	sk.AddScheduler("print", time.Second, func() { c.Add(1) })
	err := sk.Stop("print")
	if err != nil {
		t.Error(err)
	}
}

func TestIfExistStop(t *testing.T) {
	sk := New()
	c := &count{}
	interval := time.Millisecond
	sk.AddScheduler("print", interval, func() { c.Add(1) })
	time.Sleep(2 * time.Millisecond)
	if c.Get() < 1 {
		t.Fatalf("Expecting c > 0, got: %v", c.Get())
	}
	sk.AddScheduler("print", interval, func() { c.Del(1) })
	time.Sleep(3 * time.Millisecond)
	if c.Get() > 0 {
		t.Fatalf("Expecting c < 0, got: %v", c.Get())
	}
	err := sk.Stop("print")
	if err != nil {
		t.Error(err)
	}
	sk.StopAll()
}
