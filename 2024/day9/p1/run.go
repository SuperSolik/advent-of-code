package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	return string(bytes)
}

func digitsToInts(data string) []int {
	result := make([]int, len(data))
	for i, char := range data {
		result[i] = int(char - '0')
	}
	return result
}

func main() {
	inputStr := parseInput()

	digits := digitsToInts(inputStr)

	var diskMap []int

	id := 0

	for i := 0; i < len(digits); i++ {
		isNonFree := i%2 == 0

		val := -1
		if isNonFree {
			val = id
			id++
		}

		for j := 0; j < digits[i]; j++ {
			diskMap = append(diskMap, val)
		}
	}

	freeStart := 0
	fileBlocksEnd := len(diskMap) - 1

	for {
		for diskMap[freeStart] != -1 {
			freeStart++
		}

		for diskMap[fileBlocksEnd] == -1 {
			fileBlocksEnd--
		}

		if freeStart >= fileBlocksEnd {
			break
		}

		t := diskMap[fileBlocksEnd]

		diskMap[fileBlocksEnd] = diskMap[freeStart]
		diskMap[freeStart] = t

		freeStart++
		fileBlocksEnd--
	}

	result := 0

	for i := 0; i < freeStart; i++ {
		result += i * diskMap[i]
	}

	fmt.Printf("%d\n", result)
}
