package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func parseInput() []string {
	args := os.Args[1:]

	var file *os.File

	if len(args) <= 0 {
		log.Printf("Input file is not provided, reading from stdin")
		file = os.Stdin
	} else {
		inputFilePath, err := filepath.Abs(args[0])
		if err != nil {
			log.Fatalf("Unable to create abs path from: %s\n", args[0])
		}

		log.Printf("Input file: %s", inputFilePath)

		openedFile, err := os.Open(inputFilePath)

		if err != nil {
			log.Fatalf("Unable to read file: %s\n", inputFilePath)
		}

		file = openedFile
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("Unable to read content: %v\n", err)
	}

	return strings.Split(strings.Trim(string(content), " \t\n"), "\n")
}

func main() {
	lines := parseInput()

	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
}
