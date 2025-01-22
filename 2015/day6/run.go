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

func main() {
	size := 1000
	part1 := make([]uint8, size*size)
	part2 := make([]int, size*size)

	lines := parseInput()

	for _, line := range lines {
		parts := strings.Split(line, " ")

		// log.Printf("%s -> %v\n", line, parts)

		start := 0

		if len(parts) == 5 {
			start += 1
		}

		command := parts[start]

		var sx, sy, ex, ey int
		fmt.Sscanf(parts[start+1], "%d,%d", &sx, &sy)
		fmt.Sscanf(parts[start+3], "%d,%d", &ex, &ey)

		for y := sy; y <= ey; y++ {
			for x := sx; x <= ex; x++ {
				switch command {
				case "toggle":
					part1[y*size+x] ^= 1
					part2[y*size+x] += 2
				case "off":
					part1[y*size+x] = 0
					part2[y*size+x] -= 1

					if part2[y*size+x] < 0 {
						part2[y*size+x] = 0
					}
				case "on":
					part1[y*size+x] = 1
					part2[y*size+x] += 1
				}
			}
		}
		// fmt.Printf("%s, from %d,%d to %d,%d\n", command, sx, sy, ex, ey)
	}

	part1Result := 0
	part2Result := 0
	for i := 0; i < size*size; i++ {
		if part1[i] == 1 {
			part1Result++
		}

		part2Result += part2[i]
	}

	fmt.Printf("part1: %d\n", part1Result)
	fmt.Printf("part2: %d\n", part2Result)
}
