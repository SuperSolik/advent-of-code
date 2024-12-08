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

type Antenna struct {
	x, y int
}

func inBounds(x, y, w, h int) bool {
	return x >= 0 && x < w && y >= 0 && y < h
}

func main() {
	grid := parseInput()

	size, rowSize := len(grid[0]), len(grid)

	result := 0
	antennas := make(map[rune][]Antenna)

	antinodes := make([]bool, size*rowSize, size*rowSize)

	for y, row := range grid {
		for x, c := range row {
			if c == '.' {
				continue
			}
			antennas[c] = append(antennas[c], Antenna{x: x, y: y})
		}
	}

	for _, positions := range antennas {
		l := len(positions)
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				start := positions[i]
				end := positions[j]
				yDiff := end.y - start.y
				xDiff := end.x - start.x

				if inBounds(start.x-xDiff, start.y-yDiff, rowSize, size) {
					antinodes[(start.y-yDiff)*rowSize+start.x-xDiff] = true
				}

				if inBounds(end.x+xDiff, end.y+yDiff, rowSize, size) {
					antinodes[(end.y+yDiff)*rowSize+end.x+xDiff] = true
				}
			}
		}
	}

	for _, present := range antinodes {
		if present {
			result += 1
		}
	}

	fmt.Printf("%d\n", result)
}
