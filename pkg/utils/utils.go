package utils

import (
	"log"
	"os"
)

// WorkDir return working directory
func WorkDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	return path
}
