package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func writeResponse(c net.Conn, body string) {
	content := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Document</title>
</head>
<body>`
	content += body
	content += fmt.Sprintf("I see you connected from address <strong>'%v'</strong> at timestamp <strong>'%v'</strong>.<br />\n",
		c.RemoteAddr(), time.Now())
	content += `</body>
</html>
`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(content))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, content)
}
