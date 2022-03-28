package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Listener is similar to Socket in Java
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	// close socket
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!\n")

		conn.Close()
	}
}
