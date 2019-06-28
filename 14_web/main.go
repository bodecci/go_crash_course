package main

import (
	"fmt"
	"net/http" // to deal with http request
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>") // output to the browser
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	http.HandleFunc("/", index)      // kinda like a router
	http.HandleFunc("/about", about) // kinda like a router
	fmt.Println("Server starting")
	http.ListenAndServe(":3000", nil) // like app.listen in express
}
