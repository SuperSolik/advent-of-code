package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func createMatcher(patterns []string) map[byte][]string {
	m := make(map[byte][]string)

	for _, p := range patterns {
		start := p[0]
		m[start] = append(m[start], p)
	}

	return m
}

var cache = make(map[string]int)

func checkDesign(design string, matcher map[byte][]string) int {
	if ways, cached := cache[design]; cached {
		return ways
	}

	targetSize := len(design)

	if targetSize <= 0 {
		return 1
	}
	ways := 0
	if patterns, ok := matcher[design[0]]; ok {
		for _, p := range patterns {
			pSize := len(p)
			if pSize <= targetSize && strings.HasPrefix(design, p) {
				ways += checkDesign(design[pSize:], matcher)
			}
		}
	}

	cache[design] = ways
	return ways
}

func parsePatterns(line, sep string) []string {
	parts := strings.Split(strings.Trim(line, " \t\n,"), sep)

	result := make([]string, len(parts))

	for i, part := range parts {
		result[i] = strings.Trim(part, " \t\n,")
	}

	return result
}

func main() {
	lines := parseInput()
	patterns := parsePatterns(lines[0], ",")
	designs := lines[2:]

	m := createMatcher(patterns)

	result := 0
	for i, d := range designs {
		ways := checkDesign(d, m)
		fmt.Printf("%d: %s, ways = %d\n", i, d, ways)
		result += ways
	}

	fmt.Printf("%v\n", result)
}
