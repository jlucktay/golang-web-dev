package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

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
	case rMethod == "GET" && rTarget == "/favicon.ico":
		// empty branch to stop the server from crashing out
	case rMethod == "GET" && rTarget == "/":
		getRoot(c)
	case rMethod == "GET" && rTarget == "/apply":
		getApply(c)
	case rMethod == "POST" && rTarget == "/apply":
		postApply(c)
	default:
		log.Fatalf("unknown verb '%v' or path '%v'", rMethod, rTarget)
	}
}
