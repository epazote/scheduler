[![Build Status](https://travis-ci.org/epazote/scheduler.svg?branch=master)](https://travis-ci.org/epazote/scheduler)
[![Coverage Status](https://coveralls.io/repos/github/epazote/scheduler/badge.svg?branch=master)](https://coveralls.io/github/epazote/scheduler?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/epazote/scheduler)](https://goreportcard.com/report/github.com/epazote/scheduler)

Calls a function every N seconds.

Example
-----
```
// Create a Scheduler to run a function every second
s := scheduler.New()
s.AddScheduler("every second", 1, func(){
    fmt.Println("Second passed")
})

// Let the scheduler run for five seconds
time.Sleep(time.Second * 5)

// Stop the "every second" scheduler
s.Stop("every second)
```
