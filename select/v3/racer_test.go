package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(100 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// So that it does not continue to listen to a port.
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatal("did not expect an error, but got one")
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Return an error if a server doesn't respond within specified time", func(t *testing.T) {
		serverA := makeDelayedServer(1 * time.Second)

		defer serverA.Close()

		_, err := ConfigureRacer(serverA.URL, serverA.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
