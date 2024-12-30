package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func parseInput() string {
	args := os.Args[1:]

	if len(args) <= 0 {
		log.Fatal("Input file is not provided")
	}
	inputFilePath, err := filepath.Abs(args[0])
	if err != nil {
		log.Fatalf("Unable to create abs path from: %s\n", args[0])
	}

	log.Printf("Input file: %s", inputFilePath)

	bytes, err := os.ReadFile(inputFilePath)

	if err != nil {
		log.Fatalf("Unable to read file: %s\n", inputFilePath)
	}

	return strings.Trim(string(bytes), " \t\n")
}

func main() {
	parens := parseInput()

	upCnt := strings.Count(parens, "(")
	downCnt := strings.Count(parens, ")")

	fmt.Printf("part1: %v\n", upCnt-downCnt)

	floor := 0
	for i, c := range parens {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor < 0 {
			fmt.Printf("part2: %v\n", i+1)
			break
		}
	}
}
