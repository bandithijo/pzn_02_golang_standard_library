package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "rizqi,nur,assyaufi\n" +
	             "pratama,kedua,ketiga\n" +
				 "eka,dwi,tri\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
	// [rizqi nur assyaufi]
	// [pratama kedua ketiga]
	// [eka dwi tri]
}
