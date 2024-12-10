package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func parseInput() string {
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

	return string(bytes)
}

func digitsToInts(data string) []int {
	result := make([]int, len(data))
	for i, char := range data {
		result[i] = int(char - '0')
	}
	return result
}

type Block struct {
	start, size, cap int
}

func main() {
	inputStr := parseInput()

	digits := digitsToInts(inputStr)

	var diskMap []int

	id := 0

	var freeBlocks []Block
	var fileBlocks []Block

	blockStartIdx := 0

	for i := 0; i < len(digits); i++ {
		isNonFree := i%2 == 0

		val := -1
		if isNonFree {
			val = id
			id++

			fileBlocks = append(fileBlocks, Block{start: blockStartIdx, size: digits[i]})
		} else {
			freeBlocks = append(freeBlocks, Block{start: blockStartIdx, size: digits[i]})
		}

		for j := 0; j < digits[i]; j++ {
			diskMap = append(diskMap, val)
		}

		blockStartIdx += digits[i]
	}

	fileBlocksEnd := len(fileBlocks) - 1

	for fileBlocksEnd >= 0 {
		freeStart := 0

		// NOTE: trying to find the free block to move the file block to
		for freeStart < len(freeBlocks) && (freeBlocks[freeStart].size-freeBlocks[freeStart].cap) < fileBlocks[fileBlocksEnd].size {
			freeStart++
		}

		// NOTE:  unable to move the current file block
		if freeStart >= len(freeBlocks) || freeBlocks[freeStart].start > fileBlocks[fileBlocksEnd].start {
			fileBlocksEnd--
			continue
		}

		freeBlock := &freeBlocks[freeStart]
		fileBlock := &fileBlocks[fileBlocksEnd]

		// NOTE: swap free and file blocks
		temp := make([]int, fileBlock.size)
		copy(temp, diskMap[fileBlock.start:fileBlock.start+fileBlock.size])
		copy(diskMap[fileBlock.start:fileBlock.start+fileBlock.size], diskMap[freeBlock.start+freeBlock.cap:freeBlock.start+freeBlock.cap+fileBlock.size])
		copy(diskMap[freeBlock.start+freeBlock.cap:freeBlock.start+freeBlock.cap+fileBlock.size], temp)

		// NOTE: update free block cap, move to another file block
		freeBlock.cap += fileBlock.size
		fileBlocksEnd--
	}

	result := 0

	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] != -1 {
			result += i * diskMap[i]
		}
	}

	fmt.Printf("%d\n", result)
}
