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

func checkUpdate(update []int, rules map[int][]int) bool {
	size := len(update)

	idx := make(map[int]int, size)

	for i := 0; i < size; i++ {
		idx[update[i]] = i
	}

	for i, page := range update {
		nextPages, ok := rules[page]
		if !ok {
			continue
		}
		for _, nextPage := range nextPages {
			j, ok := idx[nextPage]
			if !ok {
				continue
			}

			if i > j {
				// log.Printf("Violation of rule: %d -> %d", page, nextPage)
				return false
			}
		}

	}

	return true
}

func fixUpdate(update []int, rules map[int][]int) []int {
	size := len(update)

	weight := make(map[int]int, size)

	for _, page := range update {
		nextPages, ok := rules[page]
		if !ok {
			continue
		}
		for _, nextPage := range nextPages {
			w := weight[nextPage]
			weight[nextPage] = w + 1
		}
	}

	sort.Slice(update, func(i, j int) bool {
		return weight[update[i]] < weight[update[j]]
	})

	return update
}

func main() {
	lines := parseInput()

	rules := make(map[int][]int)

	i := 0
	// NOTE: rules part
	for lines[i] != "" {

		pair, err := parseSepInts(lines[i], "|")

		if err != nil {
			log.Fatalf("Unable to parse %s with sep = %s", lines[i], "|")
		}
		values := rules[pair[0]]
		values = append(values, pair[1])
		rules[pair[0]] = values
		i++
	}

	// NOTE: skip blank line separator
	i++

	result := 0
	// NOTE: check updates part
	for i < len(lines) {

		update, err := parseSepInts(lines[i], ",")

		if err != nil {
			log.Fatalf("Unable to parse %s with sep = %s", lines[i], ",")
		}

		valid := checkUpdate(update, rules)

		if !valid {
			update = fixUpdate(update, rules)
			middle := update[len(update)/2]
			result += middle
		}

		i++
	}

	fmt.Printf("%d\n", result)
}
