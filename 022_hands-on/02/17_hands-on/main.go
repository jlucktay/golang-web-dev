// Add code to respond to the following METHODS & ROUTES:
// - GET /
// - GET /apply
// - POST /apply

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	l, errL := net.Listen("tcp", ":8080")
	if errL != nil {
		log.Fatal(errL)
	}
	defer l.Close()

	for {
		c, errC := l.Accept()
		if errC != nil {
			log.Fatal(errC)
			continue
		}
		go serve(c)
	}
}

func serve(c net.Conn) {
	defer c.Close()

	firstLine := true
	var rMethod, rTarget string
	s := bufio.NewScanner(c)

	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)

		if firstLine {
			xs := strings.Fields(ln)
			rMethod, rTarget = xs[0], xs[1]
			fmt.Println("Method:", rMethod)
			fmt.Println("Request target:", rTarget)
		}

		if ln == "" {
			fmt.Println("End of HTTP headers")
			break
		}

		firstLine = false
	}

	switch {
	case rMethod == "GET" && rTarget == "/":
		getRoot(c)
	case rMethod == "GET" && rTarget == "/favicon.ico":
		// empty branch to stop the server from crashing out
	default:
		log.Fatalf("unknown verb '%v' or path '%v'", rMethod, rTarget)
	}
}

func getRoot(c net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Document</title>
	</head>
	<body>
	`
	body += "<h1>HOLY COW THIS IS LOW LEVEL</h1>\n"
	body += fmt.Sprintf("I see you connected from address <strong>'%v'</strong> at timestamp <strong>'%v'</strong>.<br />\n",
		c.RemoteAddr(), time.Now())
	body += "Method: GET<br />\n"
	body += "Request target: /<br />\n"
	body += "</body>\n</html>"

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func getApply()  {}
func postApply() {}
