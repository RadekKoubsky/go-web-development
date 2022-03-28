package main

import (
	"fmt"
	"net/http"
)

type CarHandler struct{}

// CarHandler implements Handler interface
func (carHandler CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var carHandler CarHandler
	http.ListenAndServe(":8080", carHandler)
}
