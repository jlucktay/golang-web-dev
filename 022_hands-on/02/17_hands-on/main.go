package main

import (
	"log"
	"net"
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
