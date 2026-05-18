package main

import (
	"fmt"
	"net/http"
)

// handler function to respond to Http request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func nextHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've clicked: %s", r.URL.Path)
}

func main() {
	//registering the handler function for a specific path
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/next", nextHandler)

	//starting the server on port 8080
	fmt.Println("Starting server at port 8090")

	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Println(err)
	}
}
