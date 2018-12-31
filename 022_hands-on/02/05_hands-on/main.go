package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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

		s := bufio.NewScanner(c)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		defer c.Close()

		fmt.Println("Code got here.")
		io.WriteString(c, fmt.Sprintf("I see you connected from '%v' at '%v'.\r\n", c.RemoteAddr(), time.Now()))
	}
}
