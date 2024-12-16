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

type Robot struct {
	x, y   int
	vx, vy int
}

func (r *Robot) move(w, h, steps int) {
	r.x = (r.x + steps*r.vx) % w
	r.y = (r.y + steps*r.vy) % h

	if r.x < 0 {
		r.x = w + r.x
	}
	if r.x >= w {
		r.x = r.x - w
	}
	if r.y < 0 {
		r.y = h + r.y
	}
	if r.y >= h {
		r.y = r.y - h
	}
}

func (r *Robot) getQuadrant(w, h int) int {
	switch {
	case r.x >= 0 && r.x < w/2 && r.y >= 0 && r.y < h/2:
		return 0
	case r.x > w/2 && r.x < w && r.y >= 0 && r.y < h/2:
		return 1
	case r.x >= 0 && r.x < w/2 && r.y > h/2 && r.y < h:
		return 2
	case r.x > w/2 && r.x < w && r.y > h/2 && r.y < h:
		return 3
	default:
		return -1
	}
}

func main() {
	lines := parseInput()

	w, h := 101, 103

	var robots []Robot

	for i := 0; i < len(lines); i++ {
		robot := Robot{}
		fmt.Sscanf(lines[i], "p=%d,%d v=%d,%d", &robot.x, &robot.y, &robot.vx, &robot.vy)
		robot.move(w, h, 100)
		robots = append(robots, robot)
	}

	quadrants := make([]int, 4)

	for _, r := range robots {
		qIdx := r.getQuadrant(w, h)
		if qIdx != -1 {
			quadrants[qIdx] += 1
		}
	}

	result := 1
	for _, q := range quadrants {
		result *= q
	}

	fmt.Printf("%v\n", result)

}
