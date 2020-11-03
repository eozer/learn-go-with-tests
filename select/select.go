package main

import (
	"errors"
	"net/http"
	"time"
)

// Racer is a function
func Racer(URL1, URL2 string) (string, error) {
	select {
	case <-ping(URL1):
		return URL1, nil
	case <-ping(URL2):
		return URL2, nil
	case <-time.After(10 * time.Second):
		return "", errors.New("Timeout error")
	}
}

// NOTE: struct {} does not allocate anything in memory when make(chan struct{}) vs
// make(chan bool)
func ping(URL string) chan struct{} {
	ch := make(chan struct{}) // NOTE: always use make vs -> var ch struct {}, because ch is now nil
	go func() {
		http.Get(URL)
		close(ch)
	}()
	return ch
}

func main() {
	Racer("", "")
}
