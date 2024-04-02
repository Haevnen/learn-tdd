package racer

import "testing"

func TestRacer(t *testing.T) {
	t.Run("Compare the speed of servers, returning the url of the fastest one", func(t *testing.T) {
		slowUrl := "http://www.facebook.com"
		fastUrl := "http://www.google.com"

		want := fastUrl
		got := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
