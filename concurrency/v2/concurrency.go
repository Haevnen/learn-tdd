package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(url string) {
			// The common problem when we use closure is not capture the input
			// param
			results[url] = wc(url)
		}(url)
	}

	// Let wait 2 second for all goroutines to finish
	time.Sleep(2 * time.Second)
	return results
}
