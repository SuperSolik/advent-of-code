package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
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

	numbers := strings.Fields(string(bytes))

	var left []int
	var right []int

	for i := 0; i < len(numbers)-1; i += 2 {
		l, err := strconv.Atoi(numbers[i])
		if err != nil {
			log.Fatalf("Unable to parse %s to int", numbers[i])
		}
		r, err := strconv.Atoi(numbers[i+1])
		if err != nil {
			log.Fatalf("Unable to parse %s to int", numbers[i])
		}
		left = append(left, l)
		right = append(right, r)
	}

	if len(left) != len(right) {
		log.Fatalf("Arrays are not of the same length: %v != %v", len(left), len(right))
	}

	counts := make(map[int]int)

	for _, val := range right {
		c := counts[val]
		counts[val] = c + 1
	}

	result := 0
	for _, val := range left {
		result += val * counts[val]
	}

	fmt.Printf("%d\n", result)
}
