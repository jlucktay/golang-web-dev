package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
)

func main() {
	fmt.Println(getCode("test@example.com"))
	fmt.Println(getCode("test@exampl.com"))
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	if _, err := io.WriteString(h, s); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("% x", h.Sum(nil))
}
