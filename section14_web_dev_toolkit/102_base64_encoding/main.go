package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "String with special characters like \\ and \""

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println("Original string:")
	fmt.Println(s)
	fmt.Println("Encoded to base64:")
	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("I'm giving her all she's got Captain!", err)
	}
	fmt.Println("Decoded back from base64:")
	fmt.Println(string(bs))
}
