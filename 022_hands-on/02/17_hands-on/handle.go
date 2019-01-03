package main

import (
	"net"
)

func getRoot(c net.Conn) {
	body := `<h1>HOLY COW THIS IS LOW LEVEL</h1>
Method: GET<br />
Request target: /<br />
`

	writeResponse(c, body)
}

func getApply(c net.Conn) {
	body := `<h1>POST TO THIS FORM</h1>
Method: GET<br />
Request target: /apply<br />

<form method="POST" action="/apply">
<input type="submit" value="apply">
</form>
`

	writeResponse(c, body)
}

func postApply(c net.Conn) {
	body := `<h1>YOU POSTED TO THIS FORM</h1>
Method: POST<br />
Request target: /apply<br />
`

	writeResponse(c, body)
}
