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
func (v *Vec2) neg() Vec2 {
	return Vec2{x: -v.x, y: -v.y}
}

var DIRS = []Vec2{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

type Solution struct {
	grid  []string
	w, h  int
	pos   Vec2
	moves []byte
}

func (s *Solution) PrintGrid() {
	for y := 0; y < s.h; y++ {
		for x := 0; x < s.w; x++ {
			if x == s.pos.x && y == s.pos.y {
				fmt.Printf("@")
			} else {
				fmt.Printf("%c", s.grid[y][x])
			}
		}
		fmt.Println()
	}
}

func (s *Solution) CountGPSCoords() int {
	res := 0
	for y := 0; y < s.h; y++ {
		for x := 0; x < s.w; x++ {
			if s.grid[y][x] == 'O' {
				res += y*100 + x
			}
		}
	}
	return res
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
	if s.InBounds(pos) {
		return s.At(pos) == '#'
	}
	return true
}

func strPut(data string, idx int, c byte) string {
	row := []byte(data)
	row[idx] = c
	data = string(row)
	return data
}

func (s *Solution) loop() {
	i := 0
	for i < len(s.moves) {
		dir, ok := s.DirFromMove(s.moves[i])
		if !ok {
			i++
			continue
		}

		newPos := s.pos.plus(dir)

		// NOTE: right next to a wall
		if s.EndOfGrid(newPos) {
			i++
			continue
		}

		// NOTE: peek through boxes until either a wall or a free space
		for !s.EndOfGrid(newPos) && s.grid[newPos.y][newPos.x] != '.' {
			newPos = newPos.plus(dir)
		}

		// NOTE: wall after boxes, can't move
		if s.EndOfGrid(newPos) {
			i++
			continue
		}

		// NOTE: free space after the boxes, move boxes from the end
		backtrackPos := newPos.plus(dir.neg())
		for backtrackPos != s.pos {
			t := s.grid[backtrackPos.y][backtrackPos.x]
			s.grid[backtrackPos.y] = strPut(s.grid[backtrackPos.y], backtrackPos.x, s.grid[newPos.y][newPos.x])
			s.grid[newPos.y] = strPut(s.grid[newPos.y], newPos.x, t)

			newPos = backtrackPos
			backtrackPos = newPos.plus(dir.neg())
		}

		// NOTE: move player after the boxes are moved
		s.pos = backtrackPos.plus(dir)
		i++
	}
}

func main() {
	lines := parseInput()

	var grid []string
	i := 0

	sx, sy := -1, -1

	for ; lines[i] != ""; i++ {
		if rowIdx := strings.IndexByte(lines[i], '@'); rowIdx != -1 {
			sx = rowIdx
			sy = i
			lines[i] = strPut(lines[i], rowIdx, '.')
		}
		grid = append(grid, lines[i])
	}

	i += 1

	var moves []byte
	for ; i < len(lines); i++ {
		moves = append(moves, []byte(lines[i])...)
	}

	s := Solution{grid, len(grid[0]), len(grid), Vec2{sx, sy}, moves}

	s.PrintGrid()

	s.loop()

	fmt.Println("---AFTER---")

	s.PrintGrid()

	fmt.Printf("%v\n", s.CountGPSCoords())
}
