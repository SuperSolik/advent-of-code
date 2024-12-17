package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

func parseSepInts(line, sep string) ([]int, error) {
	parts := strings.Split(strings.Trim(line, " \t\n,"), sep)

	ints := make([]int, len(parts))

	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return ints, err
		}
		ints[i] = num
	}

	return ints, nil
}

type Computer struct {
	A, B, C int
	program []int
	out     []int
}

func (c *Computer) GetValue(op int) int {
	switch op {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return op
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		log.Fatalf("Invalid operand: %v\n", op)
	}
	return -1
}

func (c *Computer) run() {
	instructionCnt := 0

	for instructionCnt < len(c.program) {
		instruction := c.program[instructionCnt]
		instructionOp := c.program[instructionCnt+1]
		switch instruction {
		case 0: // NOTE: adv
			num := c.A
			denom := 1 << c.GetValue(instructionOp)
			c.A = num / denom
		case 1: // NOTE: bxl
			c.B = c.B ^ instructionOp
		case 2: // NOTE: bst
			c.B = c.GetValue(instructionOp) % 8
		case 3: // NOTE: jnz
			if c.A != 0 {
				instructionCnt = instructionOp
				continue
			}
		case 4: // NOTE: bxc
			c.B = c.B ^ c.C
		case 5: // NOTE: out
			c.out = append(c.out, c.GetValue(instructionOp)%8)
		case 6: // NOTE: bdv
			num := c.A
			denom := 1 << c.GetValue(instructionOp)
			c.B = num / denom
		case 7: // NOTE: cdv
			num := c.A
			denom := 1 << c.GetValue(instructionOp)
			c.C = num / denom
		default:
			fmt.Printf("Uknown instruction: %v\n", instruction)
		}

		instructionCnt += 2
	}
}

func main() {
	lines := parseInput()

	A := 0
	B := 0
	C := 0

	var programStr string

	fmt.Sscanf(lines[0], "Register A: %d", &A)
	fmt.Sscanf(lines[1], "Register B: %d", &B)
	fmt.Sscanf(lines[2], "Register C: %d", &C)

	fmt.Sscanf(lines[4], "Program: %s", &programStr)

	program, _ := parseSepInts(programStr, ",")

	fmt.Printf("%v %v %v %v\n", A, B, C, program)

	computer := Computer{A: A, B: B, C: C, program: program}

	computer.run()

	fmt.Printf("%v\n", computer)

	for i, c := range computer.out {
		fmt.Printf("%d", c)
		if i < len(computer.out)-1 {
			fmt.Print(",")
		}
	}
}
