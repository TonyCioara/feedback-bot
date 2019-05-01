package tests

import (
	"testing"
	"github.com/TonyCioara/feedback-bot/utils"
)

// TestWriteCSV tests the write CSV method from the utils package
func TestWriteCSV(t *testing.T) {
	row1 := []string{"First", "Row"}
	row2 := []string{"Second, Row"}
	rows := [][]string{row1, row2}
	err := utils.WriteCSV("testFile", rows)
	if err != nil { 
		t.Error("Failed to write CSV")
	}

	utils.DeleteFile("./" + "testFile")
}