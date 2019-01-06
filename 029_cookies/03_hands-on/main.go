package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const cookieName = "my-counter"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, errCounter := req.Cookie(cookieName)
	if errCounter == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  cookieName,
			Value: "0",
		}
	}

	counter, errConvert := strconv.Atoi(cookie.Value)
	if errConvert != nil {
		log.Fatal(errConvert)
	}
	counter++
	cookie.Value = strconv.Itoa(counter)

	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "COUNTER: '%v'\n", cookie.Value)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(cookieName)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "YOUR COUNTER: '%#v'\n", c)
}
