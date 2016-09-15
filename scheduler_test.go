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
func (c *count) Get() int64 {
	return atomic.LoadInt64(&c.i)
}

func TestScheduler(t *testing.T) {
	sk := New()
	c := &count{}
	sk.AddScheduler("print", 1, func() { fmt.Println(c.Get()); c.Add(1) })
	select {
	case <-time.After(3 * time.Second):
		sk.StopAll()
	}
	if c.Get() != 3 {
		t.Error("Expecting 3")
	}
}
