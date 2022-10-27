package main

import (
	"hart/src/scanner"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalf("You didn't specify file.")
	} else if len(args) == 1 {
		file, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatalf("File doesn't exist.")
		}
		scanner.Scanner(file)
	}
}
