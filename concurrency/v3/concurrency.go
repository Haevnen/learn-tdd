package concurrency

import "fmt"

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// let try to use channel to prevent concurrent write
	resultChan := make(chan result)

	for _, url := range urls {
		go func(url string) {
			fmt.Println("Start goroutine for", url)
			resultChan <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChan
		results[r.string] = r.bool
	}

	return results
}
