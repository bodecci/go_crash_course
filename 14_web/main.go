package main // getting the server running.

import (
	"fmt"
	"net/http" // to deal with http request
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Folks</h1>") // output to the browser
	fmt.Fprintf(w, "<h1>Going thru the gin framework</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b><center>Learning Go and building REST APIs</center></b>")
}

func standUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<b><center><h1>Standup Recap</h1></center>
					<ul><h3><li>Feel somewhat comfortable with Go.</h3></li>
					<li><h3>REST API had bugs so, just showing this off</h3></li>
					<li><h3>Very Happy I got the far.</h3></li>
					<li><h2>I'll bug Dave if I get stumped. Stack overflow-ing it for now</h2></li>
					<li><h3>Thanks for making me feel welcomed</h3></li>
					</ul></b>`)
}

func main() {
	http.HandleFunc("/", index)      // kinda like a router
	http.HandleFunc("/about", about) // kinda like a router
	http.HandleFunc("/info", info)
	http.HandleFunc("/standUp", standUp)
	fmt.Println("Server starting")
	http.ListenAndServe(":3000", nil) // like app.listen in express
}
