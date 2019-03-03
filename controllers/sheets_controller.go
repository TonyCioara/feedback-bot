package controllers

import (
	"fmt"
	"log"

	spreadsheet "gopkg.in/Iwark/spreadsheet.v2"
)

func CreateSpreadsheet(name string, rows [][]string) {

	service, err := spreadsheet.NewService()

	if err != nil {
		log.Fatalf("Error loading spreadsheet: %s", err)
	}

	fmt.Println(service)

	cellRows := [][]spreadsheet.Cell{}
	for rowIndex, row := range rows {
		cellRow := []spreadsheet.Cell{}
		for columnIndex, item := range row {
			newCell := spreadsheet.Cell{
				Row:    uint(rowIndex),
				Column: uint(columnIndex),
				Value:  item,
			}
			cellRow = append(cellRow, newCell)
		}
		cellRows = append(cellRows, cellRow)
	}

	sheet := spreadsheet.Sheet{
		Rows: cellRows,
	}
	sheets := []spreadsheet.Sheet{sheet}

	ss, err := service.CreateSpreadsheet(spreadsheet.Spreadsheet{
		Properties: spreadsheet.Properties{
			Title: name,
		},
		Sheets: sheets,
	})

	if err != nil {
		log.Fatalf("Error loading spreadsheet: %s", err)
	}

	fmt.Println("Spreadsheet:", ss)
}
