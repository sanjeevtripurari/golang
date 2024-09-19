package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":

		fmt.Fprint(w, "Hello World.. GO lang page!!")
	case "/ninja":
		fmt.Fprint(w, "Hello Ninja")
	default:
		now := time.Now()
		fmt.Fprint(w, "Current date and"+now.String())
	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc("/", htmlVsPlain)
	http.ListenAndServe("", nil)
}
