package main

import (
	"net/http"
	"time"
)

// measureHTTPResponse returns the duration took to return a response from URL.
func measureHTTPResponse(URL string) time.Duration {
	start := time.Now()
	http.Get(URL)
	duration := time.Since(start) // shorthand for: time.Now().Sub(start)
	return duration
}

// Racer is a function
func Racer(URL1, URL2 string) (fastestURL string) {
	duration1 := measureHTTPResponse(URL1)
	duration2 := measureHTTPResponse(URL2)
	// If duration1 is smaller URL1 is faster.
	if duration1 < duration2 {
		return URL1
	}
	return URL2
}

func main() {
	Racer("", "")
}
