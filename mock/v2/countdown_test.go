package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	// keep track of how many time it's called
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	// Our tests take 3 seconds to run
	// * Slow tests ruin developer productivity
	// * We have not tested an important property of our function
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}

	Countdown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if sleeper.Calls != 3 {
		t.Errorf("sleeper.Calls want 3 got %d", sleeper.Calls)
	}
}
