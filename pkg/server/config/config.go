package config

import (
	"log"
	"os"
)

var cfg = getCFG()

func init() {
	printCFG()
}

func printCFG() {
	log.Println(cfg)
}

func getCFG() string {
	path, err := os.Getwd()
	if err != nil {
		return err.Error()
	}
	return path
}
