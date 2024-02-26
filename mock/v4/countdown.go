package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// 1. If a test is working with more than 3 mocks then it is a red flag
// 2. Refactor means the code changes but the behavior stays the same
// -> favour testing expected behavior rather than the implementation

func Countdown(writer io.Writer, sleeper Sleeper) {
	// This code can pass the unittest
	// even though, the implementation is wrong

	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
