package main

import (
	"fmt"
	"log"
	"math"
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

type Equation struct {
	kXA, kYA     int
	kXB, kYB     int
	goalX, goalY int
}

// NOTE: solve by eliminating A
func (e *Equation) Solve() (int, int, bool) {
	B := float64(e.goalX*e.kYA-e.goalY*e.kXA) / float64(e.kXB*e.kYA-e.kYB*e.kXA)

	if math.Mod(B, 1) != 0 {
		return 0, 0, false
	}

	A := (float64(e.goalX) - B*float64(e.kXB)) / float64(e.kXA)

	if math.Mod(A, 1) != 0 {
		return 0, 0, false
	}

	return int(A), int(B), true
}

func main() {
	lines := parseInput()

	var kXA, kYA int
	var kXB, kYB int
	var goalX, goalY int

	result := 0
	i := 0
	for i < len(lines) {
		equationStr := lines[i : i+3]
		fmt.Sscanf(equationStr[0], "Button A: X+%d, Y+%d", &kXA, &kYA)
		fmt.Sscanf(equationStr[1], "Button B: X+%d, Y+%d", &kXB, &kYB)
		fmt.Sscanf(equationStr[2], "Prize: X=%d, Y=%d", &goalX, &goalY)

		i += 4

		// NOTE: add +10000000000000 to goalX and goalY for part 2
		equation := Equation{kXA: kXA, kXB: kXB, kYA: kYA, kYB: kYB, goalX: goalX, goalY: goalY}

		A, B, ok := equation.Solve()

		if ok {
			fmt.Printf("#%d: A:%d B:%d\n", i/4, A, B)
			result += 3*A + B

		} else {
			fmt.Printf("#%d: Not solvable\n", i/4)
		}
	}

	fmt.Printf("%v\n", result)
}
