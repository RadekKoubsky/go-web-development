package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	// we can convert func d to HandlerFunc which implements Handler interface, then we can pass it to Handle func
	// we can do the conversion because HandlerFunc has the underlying type of func(ResponseWriter, *Request)
	http.Handle("/dog", http.HandlerFunc(d))

	http.ListenAndServe(":8080", nil)
}

// this is similar to this:
// https://play.golang.org/p/X2dlgVSIrd
// ---and this---
// https://play.golang.org/p/YaUYR63b7L
