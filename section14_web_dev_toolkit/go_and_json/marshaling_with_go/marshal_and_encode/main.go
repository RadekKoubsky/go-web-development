package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// struct name does not need to be capitalized to be exported to json
type person struct {
	/*
		Fields need to be capitalized (starting with upper case) to be exported to json
	*/
	Fname string
	Lname string
	Items []string
}

// Blog about Go and Json https://go.dev/blog/json
// Convert json to a Go struct https://mholt.github.io/json-to-go/
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>FOO</title>
		</head>
		<body>
		You are at foo
		</body>
		</html>`
	w.Write([]byte(s))
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	// marshal encodes a value to json and stores it to a new variable, thus storing the json object in memory
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	// encoder writes the json data directly to writer/output stream without storing the json object in memory
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
