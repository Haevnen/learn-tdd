package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	// Our tests take 3 seconds to run
	// * Slow tests ruin developer productivity
	// * We have not tested an important property of our function
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
