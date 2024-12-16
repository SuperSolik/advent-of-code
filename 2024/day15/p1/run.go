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

type Vec2 struct {
	x, y int
}

func (v *Vec2) plus(other Vec2) Vec2 {
	return Vec2{x: v.x + other.x, y: v.y + other.y}
}

var DIRS = []Vec2{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

type Solution struct {
	grid   []string
	w, h   int
	sx, sy int
	moves  []byte
}

func (s *Solution) InBounds(pos Vec2) bool {
	return pos.x >= 0 && pos.x < s.w && pos.y >= 0 && pos.y < s.h
}

func (s *Solution) At(pos Vec2) byte {
	return s.grid[pos.y][pos.x]
}

func (s *Solution) DirFromMove(move byte) (Vec2, bool) {
	switch move {
	case '<':
		return DIRS[0], true
	case '>':
		return DIRS[1], true
	case 'v':
		return DIRS[2], true
	case '^':
		return DIRS[3], true
	default:
		return Vec2{}, false
	}
}

func (s *Solution) EndOfGrid(pos Vec2) bool {
	return !s.InBounds(pos) || s.At(pos) == '#'
}

func (s *Solution) loop() {
	pos := Vec2{s.sx, s.sy}
	for _, move := range s.moves {
		dir, ok := s.DirFromMove(move)
		if !ok {
			continue
		}

		newPos := pos.plus(dir)

		if s.EndOfGrid(newPos) {
			continue
		}

		if s.grid[newPos.y][newPos.x] == '.' {
		}

		// TODO: blocks ahead
	}
}

func main() {
	lines := parseInput()

	var grid []string
	i := 0

	sx, sy := -1, -1

	for ; lines[i] != ""; i++ {
		grid = append(grid, lines[i])
		if rowIdx := strings.IndexByte(lines[i], '@'); rowIdx != -1 {
			sx = rowIdx
			sy = i
		}
	}

	i += 1

	var moves []rune
	for ; i < len(lines); i++ {
		moves = append(moves, []rune(lines[i])...)
	}

	fmt.Printf("%d %d\n", sx, sy)

	for _, row := range grid {
		fmt.Println(row)
	}

	fmt.Println("-----")

	for _, c := range moves {
		fmt.Printf("%c", c)
	}
}
