package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	cloudPlatform = "gcp" // switch to "aws" if necessary
	connFormat    = "default:%s@tcp(%s:3306)/pocket01?charset=utf8mb4"
)

var (
	db         *sql.DB
	errOpen    error
	instanceID string

	metadataURL = map[string]string{
		"aws": "http://169.254.169.254/latest/meta-data/instance-id",
		"gcp": "http://metadata.google.internal/computeMetadata/v1/instance/id",
	}
)

func init() {
	instanceID = getInstanceID()
	connString := fmt.Sprintf(connFormat, mustEnv("MYSQL_PASSWORD"), mustEnv("MYSQL_IP"))
	db, errOpen = sql.Open("mysql", connString)
	check(errOpen)
}

func main() {
	defer db.Close()

	errPing := db.Ping()
	check(errPing)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/ping", ping)

	// http.HandleFunc("/amigos", amigos)

	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)

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
*/

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintf(w, "CREATED TABLE customer with %v rows (instance: %s)", n, instanceID)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintf(w, "INSERTED %v RECORDS (instance: %s)", n, instanceID)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	fmt.Fprintf(w, "Instance: %s\n\n", instanceID)

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintf(w, "RETRIEVED RECORD: %v\n", name)
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

	fmt.Fprintf(w, "UPDATED %v RECORDS (instance: %s)", n, instanceID)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintf(w, "DELETED %v RECORDS (instance: %s)", n, instanceID)
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
		log.Fatal(err)
	}
}

func getInstanceID() string {
	req, errReq := http.NewRequest("GET", metadataURL[cloudPlatform], nil)
	if errReq != nil {
		log.Fatal(errReq)
	}

	if cloudPlatform == "gcp" {
		req.Header.Add("Metadata-Flavor", "Google")
	}

	resp, errResp := http.DefaultClient.Do(req)
	if errResp != nil {
		log.Fatal(errResp)
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()

	return string(bs)
}

func mustEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || len(val) == 0 {
		log.Fatalf("could not find a value for key '%s' in environment", key)
	}

	return val
}
