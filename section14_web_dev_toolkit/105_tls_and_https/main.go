package main

import (
	"net/http"
)

// Use this free certificate authority https://letsencrypt.org/
// It's nonprofit open source Certificate authority provided by Internet Security Research Group
// Used by many organizations
func main() {
	http.HandleFunc("/", foo)
	// generate self-signed certificate with the following go file:
	// "go run /usr/lib/golang/src/crypto/tls/generate_cert.go --host=localhost"
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server with TLS.\n"))
}
