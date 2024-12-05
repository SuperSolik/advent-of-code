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

type Solution struct {
	lines    []string
	lineSize int
	size     int
}

func (s *Solution) LookAt(x, y int) (byte, error) {
	if x >= 0 && x < s.lineSize && y >= 0 && y < s.size {
		return s.lines[y][x], nil
	}

	return byte(0), fmt.Errorf("%d:%d out of bounds", x, y)
}

const xmas = "XMAS"

func (s *Solution) CheckXmas(x, y, dx, dy int) (bool, error) {
	bytes := make([]byte, 4, 4)
	bytes[0] = 'X'

	i := 1

	for i < len(xmas) {
		x += dx
		y += dy
		c, err := s.LookAt(x, y)
		log.Printf("%c at %d:%d\n", c, x, y)

		if err != nil {
			log.Println(err)
			return false, err
		}
		bytes[i] = c
		i += 1
	}

	res := string(bytes) == xmas

	return res, nil
}

func (s *Solution) CountXmas(x, y int) int {
	res := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			yes, _ := s.CheckXmas(x, y, dx, dy)

			if yes == true {
				fmt.Printf("Found %s at (%d, %d) with dir = (%d, %d)\n", xmas, x, y, dx, dy)
				res += 1
			}
		}
	}
	return res
}

func (s *Solution) Solve() int {
	res := 0
	for y := 0; y < s.size; y++ {
		for x := 0; x < s.lineSize; x++ {
			if s.lines[y][x] == 'X' {
				res += s.CountXmas(x, y)
			}
		}
	}
	return res
}

func main() {
	content := parseInput()

	solution := Solution{lines: content, lineSize: len(content[0]), size: len(content)}

	fmt.Printf("lineSize = %d, size = %d\n", solution.lineSize, solution.size)

	fmt.Printf("%d\n", solution.Solve())
}
