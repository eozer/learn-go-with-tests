// Package concurrency provides
package concurrency

// WebsiteChecker is a type
type WebsiteChecker func(string) bool

type WebsiteCheckerChanResult struct {
	string
	bool
}

// CheckWebsites is a function
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	wcChanRes := make(chan WebsiteCheckerChanResult)
	for _, url := range urls {
		go func(u string) {
			// results[u] = wc(u)
			wcChanRes <- WebsiteCheckerChanResult{u, wc(u)}
		}(url)
	}
	// time.Sleep(2 * time.Second)
	for i := 0; i < len(urls); i++ {
		r := <-wcChanRes
		results[r.string] = r.bool
	}

	return results
}
