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

	return strings.Split(string(bytes), " ")
}

func digitsToInts(data string) []int {
	result := make([]int, len(data))
	for i, char := range data {
		result[i] = int(char - '0')
	}
	return result
}

func runGenerations(initGen []string, genCnt int) []string {
	prevGen := initGen

	i := 0
	for i < genCnt {
		nextGen := make([]string, 0, len(prevGen)*2)
		for _, el := range prevGen {
			switch {
			case el == "0":
				nextGen = append(nextGen, "1")
			case len(el)%2 == 0:
				left, _ := strconv.Atoi(el[:len(el)/2])
				right, _ := strconv.Atoi(el[len(el)/2:])
				nextGen = append(nextGen, strconv.Itoa(left))
				nextGen = append(nextGen, strconv.Itoa(right))
			default:
				val, _ := strconv.Atoi(el)
				nextGen = append(nextGen, strconv.Itoa(val*2024))
			}
		}

		prevGen = nextGen
		i++
	}

	return prevGen
}

func main() {
	initGen := parseInput()

	finalGen := runGenerations(initGen, 75)

	fmt.Printf("%v\n", len(finalGen))
}
