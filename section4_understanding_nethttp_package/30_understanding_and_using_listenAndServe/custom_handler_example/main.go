package main

import (
	"log"
	"net/http"
)

type CarHandler struct{}

// CarHandler implements Handler interface
func (carHandler CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Any code you want in this func")
}

func main() {

}
