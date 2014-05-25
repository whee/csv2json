// Copyright (c) 2014 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license which can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	r := csv.NewReader(os.Stdin)

	header, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	record := make(map[string]string)
	enc := json.NewEncoder(os.Stdout)

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for i, field := range header {
			record[field] = row[i]
		}
		if err := enc.Encode(&record); err != nil {
			log.Fatal(err)
		}
	}
}
