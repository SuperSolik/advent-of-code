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

func (s *Solution) CheckTarget(x, y, dx, dy int, target string) (bool, error) {
	l := len(target)
	bytes := make([]byte, l, l)

	i := 0

	for i < l {
		c, err := s.LookAt(x, y)

		if err != nil {
			log.Println(err)
			return false, err
		}
		log.Printf("%c at %d:%d\n", c, x, y)
		bytes[i] = c

		x += dx
		y += dy
		i += 1
	}

	res := string(bytes) == target

	return res, nil
}

func (s *Solution) CheckXmas(x, y int) bool {

	mas1, _ := s.CheckTarget(x-1, y-1, 1, 1, "MAS")
	sam1, _ := s.CheckTarget(x-1, y-1, 1, 1, "SAM")
	mas2, _ := s.CheckTarget(x+1, y-1, -1, 1, "MAS")
	sam2, _ := s.CheckTarget(x+1, y-1, -1, 1, "SAM")

	if (mas1 || sam1) && (mas2 || sam2) {
		return true
	}

	return false
}

func (s *Solution) Solve() int {
	res := 0
	for y := 1; y < s.size-1; y++ {
		for x := 1; x < s.lineSize-1; x++ {
			if s.lines[y][x] == 'A' && s.CheckXmas(x, y) {
				res += 1
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
