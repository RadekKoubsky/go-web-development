package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.

		As a special case, the returned file server redirects any request ending in "/index.html" to the same path,
		without the final "index.html"

		If we put "index.html" to the file system root, the file server will show /index.html when accessing root path "/"
		instead of all dirs and files in root
	*/
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
