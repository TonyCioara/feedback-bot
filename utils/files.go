package utils

import (
	"fmt"
	"log"
	"os"
)

// DeleteFile deletes a file
func DeleteFile(path string) {
	// delete file
	err := os.Remove(path)
	if err != nil {
		log.Fatalf("Error de,eting file with path '%s': %s\n", path, err)
	}

	fmt.Println("==> done deleting file")
}
