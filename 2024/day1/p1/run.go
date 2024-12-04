package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
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

	sort.Ints(left)
	sort.Ints(right)

	dist := 0

	for i := 0; i < len(left); i++ {
		v := left[i] - right[i]
		if v < 0 {
			v *= -1
		}
		dist += v
	}

	fmt.Printf("%d\n", dist)
}
