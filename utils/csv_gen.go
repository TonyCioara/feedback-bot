package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCSV(name string, rows [][]string) {

	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatalf("Cannot close '%s': %s\n", name, e.Error())
		}
	}()

	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
}
