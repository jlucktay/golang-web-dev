package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type tableRecord struct {
	Date   time.Time
	Volume int

	// Open, High, Low, Close, AdjClose float64
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	r := parseRecords("table.csv")
	tpl := template.Must(template.ParseFiles("09.gohtml"))

	err := tpl.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}
}

func parseRecords(filename string) []tableRecord {
	source, errOpen := os.Open(filename)
	if errOpen != nil {
		log.Fatal(errOpen)
	}
	defer source.Close()

	tableReader := csv.NewReader(source)
	tableRecords, tableReadErr := tableReader.ReadAll()
	if tableReadErr != nil {
		log.Fatal(tableReadErr)
	}

	records := make([]tableRecord, 0, len(tableRecords))

	for i, tr := range tableRecords {
		if i == 0 {
			continue
		}

		newDate, errDateParse := time.Parse("2006-01-02", tr[0])
		if errDateParse != nil {
			log.Fatal(errDateParse)
		}

		newVolume, errVolumeParse := strconv.Atoi(tr[5])
		if errVolumeParse != nil {
			log.Fatal(errVolumeParse)
		}

		records = append(records, tableRecord{newDate, newVolume})
	}

	return records
}
