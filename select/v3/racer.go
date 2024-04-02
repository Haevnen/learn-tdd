package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// Why are we testing the speeds of the websites one after another when Go
// support Goroutine.
func ConfigureRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func Racer(a, b string) (string, error) {
	// select allow you to wait on multiple channel
	// it different with val := <-ch
	// because the latter is blocking call, as you're waiting for a value
	return ConfigureRacer(a, b, tenSecondTimeout)
}

// Why struct{}?
// Because struct{} is the smallest data type available from a memory
// perspective (so we get no allocation versus a bool)
func ping(url string) chan struct{} {
	// if we use var declaration, its default value is nil
	// if you try and send to it with <-, it will block forever
	// Because you can not send to nil channel
	// -> Always make channel
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		// Send signal to channel
		close(ch)
	}()

	return ch
}
