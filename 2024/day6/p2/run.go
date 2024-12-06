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

type Visit struct {
	dx, dy  int
	visited bool
}

func checkLoop(grid []string, rowSize, size, x, y int) bool {
	dx, dy := 0, -1

	visited := make([]Visit, rowSize*size, rowSize*size)

	for x+dx >= 0 && x+dx < rowSize && y+dy >= 0 && y+dy < size {
		visit := visited[y*rowSize+x]

		if !visit.visited {
			visit.dx = dx
			visit.dy = dy
			visit.visited = true
		} else {
			if visit.dx == dx && visit.dy == dy {
				return true
			}
		}

		visited[y*rowSize+x] = visit

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

	return false
}

func main() {
	grid := parseInput()

	sx, sy := -1, -1
	for i, row := range grid {
		if j := strings.IndexByte(row, '^'); j != -1 {
			sx, sy = j, i
			break
		}
	}

	if sx == -1 {
		log.Fatalf("Unable to find start ^ in grid")
	}
	rowSize := len(grid[0])
	size := len(grid)

	result := 0
	for y := 0; y < size; y++ {
		rowOrig := grid[y]
		for x := 0; x < rowSize; x++ {
			if (x == sx && y == sy) || grid[y][x] == '#' {
				continue
			}

			grid[y] = strPut(grid[y], x, '#')

			if loop := checkLoop(grid, rowSize, size, sx, sy); loop {
				result += 1
			}

			grid[y] = rowOrig
		}
	}

	fmt.Printf("%d\n", result)
}
