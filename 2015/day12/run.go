package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func parseInput() []byte {
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

	return content
}

func solve(data interface{}, part2 bool) float64 {
	result := 0.0
	switch v := data.(type) {
	case map[string]interface{}:
		for _, val := range v {
			if part2 {
				if strVal, ok := val.(string); ok && strVal == "red" {
					return 0.0
				}
			}
			result += solve(val, part2)
		}
	case []interface{}:
		for _, el := range v {
			result += solve(el, part2)
		}
	case float64:
		result += v
	}
	return result
}

func main() {
	jsonData := parseInput()

	// Decode into an empty interface
	var data interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("part1: %d\n", int(solve(data, false)))
	fmt.Printf("part2: %d\n", int(solve(data, true)))
}
