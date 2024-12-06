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

func strPut(data string, idx int, c rune) string {
	row := []rune(data)
	row[idx] = c
	data = string(row)
	return data
}

func main() {
	grid := parseInput()

	x, y := -1, -1
	for i, row := range grid {
		if j := strings.IndexByte(row, '^'); j != -1 {
			x, y = j, i
			break
		}
	}

	if x == -1 {
		log.Fatalf("Unable to find start ^ in grid")
	}

	dx, dy := 0, -1

	rowSize := len(grid[0])
	size := len(grid)

	for x+dx >= 0 && x+dx < rowSize && y+dy >= 0 && y+dy < size {

		grid[y] = strPut(grid[y], x, 'X')

		// NOTE: obstacles can be next to each other
		// so need to rotate until we have a clear path
		for grid[y+dy][x+dx] == '#' {
			// NOTE: turn right:
			// 0, -1 -> 1, 0
			// 1, 0 -> 0, 1
			// 0, 1 -> -1, 0
			// -1, 0 -> 0, -1
			t := dx
			dx = -dy
			dy = t
		}

		x += dx
		y += dy
	}

	grid[y] = strPut(grid[y], x, 'X')

	result := 0
	for _, row := range grid {
		result += strings.Count(row, "X")
	}

	fmt.Printf("%d\n", result)
}
