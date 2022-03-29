package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test@exampl.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
