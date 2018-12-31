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

	i := 0
	s := bufio.NewScanner(c)
	var rMethod, rTarget, rVersion string
	for s.Scan() {
		ln := s.Text()

		if i == 0 {
			i++
			rLine := strings.Fields(ln)
			rMethod, rTarget, rVersion = rLine[0], rLine[1], rLine[2]
			fmt.Println("Method:", rMethod)
			fmt.Println("Request target:", rTarget)
			fmt.Println("HTTP version:", rVersion)
		}

		if ln == "" {
			fmt.Println("End of HTTP headers")
			break
		}

		fmt.Println(ln)
	}

	body := fmt.Sprintf("I see you connected from address '%v' at timestamp '%v'.\r\n",
		c.RemoteAddr(), time.Now())
	body += fmt.Sprintf("Method: %v\r\n", rMethod)
	body += fmt.Sprintf("Request target: %v\r\n", rTarget)
	body += fmt.Sprintf("HTTP version: %v\r\n", rVersion)

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
