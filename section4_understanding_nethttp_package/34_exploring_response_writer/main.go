package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	// var initializes "d" to its zero value, thus we can pass "d" variable to ListenAndServe and
	// do not get null pointer
	var d hotdog
	http.ListenAndServe(":8080", d)
}
