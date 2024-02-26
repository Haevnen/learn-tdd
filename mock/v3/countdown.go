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

type DefaultSleeper struct{}

func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	// This code can pass the unittest
	// even though, the implementation is wrong
	/*
		for i := 3; i > 0; i-- {
			sleeper.Sleep()
		}

		for i := 3; i > 0; i-- {
			fmt.Fprintln(writer, i)
		}

		fmt.Fprint(writer, "Go!")
	*/

	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)

		// add 1-second pause
		// we have a dependency on Sleeping
		// which we need to extract, so we can then control it in our test

		// if we can mock time.Sleep, we can use dependency injection to use it
		// instead of a real time.Sleep
		// time.Sleep(1 * time.Second)

		sleeper.Sleep()
	}
	fmt.Fprint(writer, "Go!")
}

func main() {
	Countdown(os.Stdout, DefaultSleeper{})
}
