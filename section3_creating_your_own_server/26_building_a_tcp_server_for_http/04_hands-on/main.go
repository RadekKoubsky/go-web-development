package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type RequestLine struct {
	method string
	uri    string
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// parseRequest
	requestLine := parseRequest(conn)

	// write response
	respond(conn, requestLine)
}

func parseRequest(conn net.Conn) RequestLine {
	i := 0
	scanner := bufio.NewScanner(conn)
	var requestLine RequestLine
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			fields := strings.Fields(ln)
			requestLine = RequestLine{
				method: fields[0],
				uri:    fields[1],
			}
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return requestLine
}

func respond(conn net.Conn, requestLine RequestLine) {
	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>

	<a href="/">welcome page</a><br>
	<a href="/car">Get Car</a><br>
	<form method="POST" action="/car">
	<input type="submit" value="create car">
	</form>
	<strong>%v</strong><br>
	</body></html>`, processRequest(requestLine))

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func processRequest(line RequestLine) string {
	var result string
	// multiplexer, decides what to run based on URI (route) and http method
	switch line.uri {
	case "/car":
		switch line.method {
		case "POST":
			result = "Car Tesla Model S created"
		case "GET":
			result = "Car returned: Tesla Model S"
		}
	default:
		result = "Welcome page"
	}
	return result
}
