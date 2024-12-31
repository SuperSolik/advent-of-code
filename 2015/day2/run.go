package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
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

	return strings.Split(strings.Trim(string(bytes), " \t\n"), "\n")
}

func boxArea(l, w, h int) int {
	sides := []int{l * w, w * h, h * l}

	sum := sides[0] + sides[1] + sides[2]

	return slices.Min(sides) + 2*sum
}

func boxRibbon(l, w, h int) int {
	halfPerims := []int{l + w, w + h, h + l}

	return 2*slices.Min(halfPerims) + l*w*h
}

func main() {
	lines := parseInput()

	result1 := 0
	result2 := 0
	l, w, h := 0, 0, 0
	for _, line := range lines {
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		area := boxArea(l, w, h)
		ribbon := boxRibbon(l, w, h)

		fmt.Printf("%s: %d %d\n", line, area, ribbon)

		result1 += area
		result2 += ribbon
	}

	fmt.Printf("part1: %d\n", result1)
	fmt.Printf("part2: %d\n", result2)
}
