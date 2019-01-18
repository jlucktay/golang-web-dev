package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	// _ "github.com/go-sql-driver/mysql"
)

/*
var db *sql.DB
var err error
*/
var instanceID string

func init() {
	/*
		resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
		if err != nil {
			log.Fatal(err)
			return
		}

		bs := make([]byte, resp.ContentLength)
		resp.Body.Read(bs)
		resp.Body.Close()
	*/
	instanceID = "string(bs)"
}

func main() {
	/*
		db, err = sql.Open("mysql", "awsuser:mypassword@tcp(mydbinstance.cakwl95bxza0.us-west-1.rds.amazonaws.com:3306)/test02?charset=utf8")
		check(err)
		defer db.Close()

		err = db.Ping()
		check(err)
	*/
	// barebones server to start with
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/ping", ping)
	/*
		http.HandleFunc("/amigos", amigos)
		http.HandleFunc("/create", create)
		http.HandleFunc("/insert", insert)
		http.HandleFunc("/read", read)
		http.HandleFunc("/update", update)
		http.HandleFunc("/delete", del)
		http.HandleFunc("/drop", drop)
	*/
	log.Fatal(http.ListenAndServe(":80", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from GCP. (%s)", instanceID)
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

/*
func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = fmt.Sprintf("INSTANCE '%s' RETRIEVED RECORDS:\n", instanceID)

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
*/
