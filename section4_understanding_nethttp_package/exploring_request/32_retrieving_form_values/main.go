package main

import (
	"html/template"
	"log"
	"net/http"
)

type CarHandler struct{}

// CarHandler implements Handler interface
func (carHandler CarHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", request.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	// var initializes carHandler to its zero value, thus we can pass carHandler variable to ListenAndServe and
	// do not get null pointer
	var carHandler CarHandler
	http.ListenAndServe(":8080", carHandler)
}
