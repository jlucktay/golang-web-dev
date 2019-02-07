package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/authenticate", auth)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		c.Value = e + `|` + getCode(e)
	}

	http.SetCookie(w, c)

	if _, err := io.WriteString(w, `<!DOCTYPE html>
<html>
	<body>
		<form method="POST">
			<input type="email" name="email">
			<input type="submit">
		</form>
		<a href="/authenticate">Validate This `+c.Value+`</a>
	</body>
</html>`); err != nil {
		log.Fatal(err)
	}

}

func auth(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	xs := strings.Split(c.Value, "|")
	email := xs[0]
	codeRcvd := xs[1]
	codeCheck := getCode(email)

	if codeRcvd != codeCheck {
		fmt.Println("HMAC codes didn't match")
		fmt.Println(codeRcvd)
		fmt.Println(codeCheck)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if _, err := io.WriteString(w, `<!DOCTYPE html>
<html>
	<body>
		<h1>`+codeRcvd+` - RECEIVED </h1>
		<h1>`+codeCheck+` - RECALCULATED </h1>
	</body>
</html>`); err != nil {
		log.Fatal(err)
	}
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	if _, err := io.WriteString(h, data); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
