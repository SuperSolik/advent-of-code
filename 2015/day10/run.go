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

type NumberCnt struct {
	number int
	cnt    int
}

var cache = make(map[NumberCnt][]int)

func solve(numbers []int) []int {
	var newNums []int

	cnt := 1

	i := 0
	for ; i < len(numbers)-1; i++ {
		if numbers[i] != numbers[i+1] {
			newNums = append(newNums, cnt)
			newNums = append(newNums, numbers[i])
			cnt = 1
		} else {
			cnt += 1
		}
	}
	newNums = append(newNums, cnt)
	newNums = append(newNums, numbers[i])
	return newNums
}

func main() {
	input := parseInput()[0]
	numbers := make([]int, len(input))
	for i, c := range input {
		numbers[i] = int(c - '0')
	}

	for i := 0; i < 40; i++ {
		numbers = solve(numbers)
	}

	fmt.Printf("part1: %v\n", len(numbers))

	for i := 0; i < 10; i++ {
		numbers = solve(numbers)
	}

	fmt.Printf("part2: %v\n", len(numbers))
}
