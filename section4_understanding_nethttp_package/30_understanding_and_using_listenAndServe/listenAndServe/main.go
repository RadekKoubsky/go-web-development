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
	// var initializes carHandler to its zero value, thus we can pass carHandler variable to ListenAndServe and
	// do not get null pointer
	// or we can use new function to initialize the variable:
	// var carHandler = new(CarHandler)
	// or
	// carHandler := new(CarHandler)
	var carHandler CarHandler
	http.ListenAndServe(":8080", carHandler)
}
