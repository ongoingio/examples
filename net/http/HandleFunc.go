// Package: net/http
// Function: HandleFunc()
// Docs: http://golang.org/pkg/net/http/#HandleFunc
// Play: http://play.golang.org/p/9g4LgH19Ro

// HandleFunc registers a handler function for the given URL pattern.

package main

import (
	"io"
	"net/http"
)

func GreetHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, Gopher!\n")
}

func main() {
	// Registers our GreetHandler function for the /greet URL pattern. This
	// means that whenever a request reaches /greet, the HTTP server responds
	// with our GreetHandler function.
	http.HandleFunc("/greet", GreetHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
