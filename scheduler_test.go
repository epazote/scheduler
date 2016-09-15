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

func (c *count) Add() {
	atomic.AddInt64(&c.i, 1)
}

func TestScheduler(t *testing.T) {
	sk := New()
	c := &count{}
	sk.AddScheduler("print", 1, func() { fmt.Println(atomic.LoadInt64(&c.i)); c.Add() })
	select {
	case <-time.After(3 * time.Second):
		sk.StopAll()
	}
	if atomic.LoadInt64(&c.i) != 3 {
		t.Error("Expecting 3")
	}
}
