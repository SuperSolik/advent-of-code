package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func parseInput() []string {
	args := os.Args[1:]

	var file *os.File

	if len(args) <= 0 {
		log.Printf("Input file is not provided, reading from stdin")
		file = os.Stdin
	} else {
		inputFilePath, err := filepath.Abs(args[0])
		if err != nil {
			log.Fatalf("Unable to create abs path from: %s\n", args[0])
		}

		log.Printf("Input file: %s", inputFilePath)

		openedFile, err := os.Open(inputFilePath)

		if err != nil {
			log.Fatalf("Unable to read file: %s\n", inputFilePath)
		}

		file = openedFile
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("Unable to read content: %v\n", err)
	}

	return strings.Split(strings.Trim(string(content), " \t\n"), "\n")
}

type Vec2 struct {
	x, y int
}

func (v *Vec2) plus(other Vec2) Vec2 {
	return Vec2{v.x + other.x, v.y + other.y}
}

var DIRS = map[rune]Vec2{
	'>': {1, 0},
	'<': {-1, 0},
	'^': {0, -1},
	'v': {0, 1},
}

func part1(directions string) int {
	visited := make(map[Vec2]bool)

	pos := Vec2{0, 0}
	visited[pos] = true

	for _, c := range directions {
		if dir, ok := DIRS[c]; ok {
			pos = pos.plus(dir)
			visited[pos] = true
		}
	}
	return len(visited)
}

func part2(directions string) int {
	visited := make(map[Vec2]bool)

	pos := Vec2{0, 0}
	roboPos := Vec2{0, 0}

	visited[pos] = true

	for i, c := range directions {
		if dir, ok := DIRS[c]; ok {
			if i%2 == 0 {
				pos = pos.plus(dir)
				visited[pos] = true
			} else {
				roboPos = roboPos.plus(dir)
				visited[roboPos] = true
			}
		}
	}
	return len(visited)
}

func main() {
	directions := parseInput()[0]
	fmt.Printf("part1: %v\n", part1(directions))
	fmt.Printf("part2: %v\n", part2(directions))
}
