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

	return strings.Split(strings.Trim(string(bytes), " \t\n"), " ")
}

type GenCacheKey struct {
	number string
	genCnt int
}

var genCache = make(map[GenCacheKey]int)

func runGenerations(stone string, genCnt int) int {
	cacheKey := GenCacheKey{number: stone, genCnt: genCnt}

	if cnt, exists := genCache[cacheKey]; exists {
		return cnt
	}

	if genCnt <= 0 {
		return 1
	}

	if stone == "0" {
		cnt := runGenerations("1", genCnt-1)
		genCache[cacheKey] = cnt
		return cnt
	}

	if len(stone)%2 == 0 {
		left, _ := strconv.Atoi(stone[:len(stone)/2])
		right, _ := strconv.Atoi(stone[len(stone)/2:])
		cnt := runGenerations(strconv.Itoa(left), genCnt-1) + runGenerations(strconv.Itoa(right), genCnt-1)
		genCache[cacheKey] = cnt
		return cnt
	}

	val, _ := strconv.Atoi(stone)
	cnt := runGenerations(strconv.Itoa(val*2024), genCnt-1)
	genCache[cacheKey] = cnt
	return cnt
}

func main() {
	stones := parseInput()

	genCnt := 75 // 25 for part 1
	result := 0
	for _, stone := range stones {
		result += runGenerations(stone, genCnt)
	}

	fmt.Printf("%v\n", result)
}
