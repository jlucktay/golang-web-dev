package main

import (
	"encoding/csv"
	"fmt"
	"log"
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
	r := parseRecords("table.csv")

	for _, x := range r {
		fmt.Printf("%s, %#v\n", x.Date, x.Volume)
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

		newRecord := tableRecord{
			Date:   newDate,
			Volume: newVolume,
		}

		records = append(records, newRecord)
	}

	return records
}
