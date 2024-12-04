package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
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

func printSlice[T any](arr []T, newline bool) {
	for _, v := range arr {
		fmt.Printf("%v ", v)
	}
	if newline {
		fmt.Println()
	}
}

type Parser struct {
	data    string
	dataLen int
	curPos  int
}

func (p *Parser) ParseComma() error {
	if p.curPos >= p.dataLen {
		return fmt.Errorf("End of string")
	}

	c := ','

	val := rune(p.data[p.curPos])

	if val == c {
		return nil
	}
	return fmt.Errorf("Unable to parse `%c` data[%d] = %c", c, p.curPos, val)
}
func (p *Parser) ParseParen(reverse bool) error {
	if p.curPos >= p.dataLen {
		return fmt.Errorf("End of string")
	}

	c := '('
	if reverse == true {
		c = ')'
	}

	val := rune(p.data[p.curPos])

	if val == c {
		return nil
	}
	return fmt.Errorf("Unable to parse `%c` data[%d] = %c", c, p.curPos, val)
}

func (p *Parser) ParseInt() (int, int, error) {
	if p.curPos >= p.dataLen {
		return 0, 0, fmt.Errorf("End of string")
	}

	start := p.curPos
	end := start
	for unicode.IsDigit(rune(p.data[end])) && end < p.dataLen && end-start <= 3 {
		end += 1
	}

	if start == end {
		return 0, 0, fmt.Errorf("No int value staring at %d", start)
	}

	number, err := strconv.Atoi(p.data[start:end])

	return number, end - start, err
}

func (p *Parser) ParseMul() error {
	if p.curPos >= p.dataLen {
		return fmt.Errorf("End of string")
	}

	target := "mul"

	if p.curPos >= p.dataLen-1-len(target) {
		return fmt.Errorf("Not enough characters left for `%s` at %d, total size = %d", target, p.curPos, p.dataLen)
	}

	if p.data[p.curPos:p.curPos+3] == target {
		return nil
	}

	return fmt.Errorf("No `%s` starting at %d", target, p.curPos)
}

func main() {
	content := parseInput()

	parser := Parser{data: content, dataLen: len(content), curPos: 0}

	result := 0

	matches := 0

	for parser.curPos < parser.dataLen-1 {
		err := parser.ParseMul()

		if err != nil {
			log.Println(err)
			parser.curPos += 1
			continue
		} else {
			log.Printf("Matched `mul` at %d:%d\n", parser.curPos, parser.curPos+3)
			parser.curPos += 3
		}

		err = parser.ParseParen(false)
		parser.curPos += 1

		if err != nil {
			log.Println(err)
			continue
		} else {
			log.Printf("Matched `(` at %d\n", parser.curPos-1)
		}

		left, consumed, err := parser.ParseInt()

		if err != nil {
			log.Println(err)
			parser.curPos += 1
			continue
		} else {
			log.Printf("Matched %d at %d:%d\n", left, parser.curPos, parser.curPos+consumed)
			parser.curPos += consumed
		}

		err = parser.ParseComma()
		parser.curPos += 1

		if err != nil {
			log.Println(err)
			continue
		} else {
			log.Printf("Matched , at %d\n", parser.curPos-1)
		}

		right, consumed, err := parser.ParseInt()

		if err != nil {
			log.Println(err)
			parser.curPos += 1
			continue
		} else {
			log.Printf("Matched %d at %d:%d\n", right, parser.curPos, parser.curPos+consumed)
			parser.curPos += consumed
		}

		err = parser.ParseParen(true)
		parser.curPos += 1

		if err != nil {
			log.Println(err)
			continue
		} else {
			log.Printf("Matched `)` at %d\n", parser.curPos-1)
		}

		log.Printf("Matched expression `mul(%d,%d)`", left, right)

		matches += 1

		result += left * right
	}

	fmt.Printf("Result = %d, matches = %d\n", result, matches)
}
