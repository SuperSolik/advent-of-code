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

	return strings.Split(string(bytes), "\n")
}

func parseSepInts(line, sep string) ([]int, error) {
	parts := strings.Split(line, sep)

	ints := make([]int, len(parts))

	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return ints, err
		}
		ints[i] = num
	}

	return ints, nil
}

func parseEquation(line string) (int, []int) {
	parsed := strings.Split(line, ":")
	goal, numbers := parsed[0], parsed[1]

	numbers = strings.Trim(numbers, " ")

	goalParsed, err := strconv.Atoi(goal)

	if err != nil {
		log.Fatalf("Unable to parse %s to int\n", goal)
	}

	numbersParsed, err := parseSepInts(numbers, " ")

	if err != nil {
		log.Fatalf("Unable to parse %s to []int\n", numbers)
	}

	return goalParsed, numbersParsed

}

func eqSolvable(goal int, acc int, numbers []int) bool {
	if len(numbers) <= 0 {
		return acc == goal
	}
	// NOTE: try + branch
	if eqSolvable(goal, acc+numbers[0], numbers[1:]) {
		return true
	}

	// NOTE: try * branch
	if eqSolvable(goal, acc*numbers[0], numbers[1:]) {
		return true
	}

	// NOTE: try || branch

	concat := fmt.Sprintf("%d%d", acc, numbers[0])

	concatNumber, err := strconv.Atoi(concat)

	if err != nil {
		log.Fatalf("Unable to parse concatenated number %s\n", concat)
	}

	if eqSolvable(goal, concatNumber, numbers[1:]) {
		return true
	}

	return false
}

func main() {
	lines := parseInput()

	result := 0
	for _, line := range lines {
		goal, numbers := parseEquation(line)
		if eqSolvable(goal, numbers[0], numbers[1:]) {
			result += goal
		}
	}

	fmt.Printf("%d\n", result)
}
