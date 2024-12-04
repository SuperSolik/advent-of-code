package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseInput() []string {
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

	lines := strings.Split(string(bytes), "\n")
	return lines
}
func checkReport(numbers []int, size int) int {
	diff := make([]int, size-1)

	pos, neg, zeroes, max_diff := 0, 0, 0, 0

	for i := 0; i < size-1; i++ {
		diff[i] = numbers[i+1] - numbers[i]
		if diff[i] > 0 {
			max_diff = max(max_diff, diff[i])
			pos += 1
		} else if diff[i] < 0 {
			max_diff = max(max_diff, -diff[i])
			neg += 1
		} else {
			zeroes += 1
		}
	}

	if max_diff <= 3 && (neg == size-1 || pos == size-1) {
		// fmt.Printf(" safe\n")
		return 1
	}
	// fmt.Printf(" unsafe\n")
	return 0
}

func main() {
	lines := parseInput()

	result := 0

	for _, line := range lines {
		// Split the input string into a slice of strings
		lineSplit := strings.Fields(line)

		// Create a slice to hold the integers
		lineLen := len(lineSplit)
		numbers := make([]int, 0, lineLen)

		// Convert each string to an integer
		for _, str := range lineSplit {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatalf("Error converting string to int: %v", err)
			}
			numbers = append(numbers, num)
		}

		if checkReport(numbers, lineLen) == 1 {
			result += 1
		} else {
			for j := 0; j < lineLen; j++ {
				var t []int
				t = append(t, numbers[:j]...)
				t = append(t, numbers[j+1:]...)
				if checkReport(t, lineLen-1) == 1 {
					result += 1
					break
				}
			}
		}
	}

	fmt.Printf("%d\n", result)
}
