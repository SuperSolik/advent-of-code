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

func digitsToInts(data string) []int {
	result := make([]int, len(data))
	for i, char := range data {
		result[i] = int(char - '0')
	}
	return result
}

type Vec2 struct {
	x, y int
}

func (v *Vec2) plus(other Vec2) Vec2 {
	return Vec2{x: v.x + other.x, y: v.y + other.y}
}

var DIRS = []Vec2{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

type TrailMap struct {
	grid [][]int
	w, h int
}

func (t *TrailMap) inBounds(x, y int) bool {
	return x >= 0 && x < t.w && y >= 0 && y < t.h
}

func (t *TrailMap) countTrails(pos Vec2, visited map[Vec2]bool) int {

	posVal := t.grid[pos.y][pos.x]

	if posVal == 9 {
		return 1
	}

	visited[pos] = true

	res := 0

	for _, dir := range DIRS {
		neighbor := pos.plus(dir)
		if t.inBounds(neighbor.x, neighbor.y) && t.grid[neighbor.y][neighbor.x] == posVal+1 {
			if yes := visited[neighbor]; !yes {
				res += t.countTrails(neighbor, visited)
			}
		}
	}

	visited[pos] = false

	return res
}

func main() {
	lines := parseInput()
	grid := make([][]int, len(lines))

	var startPositions []Vec2

	for y, line := range lines {
		row := digitsToInts(line)

		for x, height := range row {
			if height == 0 {
				startPositions = append(startPositions, Vec2{x: x, y: y})
			}
		}
		grid[y] = row
	}

	trailMap := TrailMap{grid: grid, w: len(grid[0]), h: len(grid)}

	res := 0

	for _, startPos := range startPositions {
		visited := make(map[Vec2]bool)
		res += trailMap.countTrails(startPos, visited)
	}

	fmt.Printf("%d\n", res)
}
