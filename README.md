[![Build Status](https://travis-ci.org/epazote/scheduler.svg?branch=master)](https://travis-ci.org/epazote/scheduler)
[![Coverage Status](https://coveralls.io/repos/github/epazote/scheduler/badge.svg?branch=master)](https://coveralls.io/github/epazote/scheduler?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/epazote/scheduler)](https://goreportcard.com/report/github.com/epazote/scheduler)

Calls a function every N `time.Duration`

Example
-------

```
package main

import (
        "fmt"
        "time"

        "github.com/epazote/scheduler"
)

func main() {
        // Create new scheduler
        s := scheduler.New()

        every := time.Second
        // Add a scheduled function
        s.AddScheduler("every second", every, func() {
                fmt.Println("Second passed")
        })

        // Let scheduler run for five seconds
        time.Sleep(5 * time.Second)

        // Stop the scheduled "every second" function
        err := s.Stop("every second")
        if err != nil {
                panic(err)
        }

        // Scheduler has now stopped
        time.Sleep(5 * time.Second)
}
```
