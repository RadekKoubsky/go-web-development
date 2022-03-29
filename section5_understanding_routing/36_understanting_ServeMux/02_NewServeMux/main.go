package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	// the pattern with leading "/" will match all paths after "/dog/" to the dog handler
	mux.Handle("/dog/", d)
	// here we do not have leading "/", so the cat handler matches only "/cat" path
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
