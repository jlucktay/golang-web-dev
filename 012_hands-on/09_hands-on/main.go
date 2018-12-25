package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

type tableRecord struct {
	Date   time.Time
	Volume int

	// Open, High, Low, Close, AdjClose float64
}

func main() {
	tableContent, tableErr := ioutil.ReadFile("table.csv")
	if tableErr != nil {
		log.Fatal(tableErr)
	}

	tableReader := csv.NewReader(bytes.NewReader(tableContent))

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

	for _, x := range records {
		fmt.Printf("%s, %#v\n", x.Date, x.Volume)
	}
}
