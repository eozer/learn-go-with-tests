package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet prints
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}

// GreetHandler is a handler function which returns a message to a http response.
func GreetHandler(writer http.ResponseWriter, request *http.Request) {
	Greet(writer, "world!")
}

func main() {
	fmt.Println("Starting server at: localhost:5000")
	http.ListenAndServe(":5000", http.HandlerFunc(GreetHandler))
}
