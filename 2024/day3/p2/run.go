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

func (p *Parser) ParseText(target string) (int, error) {
	if p.curPos >= p.dataLen {
		return 0, fmt.Errorf("End of string")
	}

	l := len(target)

	if p.curPos >= p.dataLen-1-l {
		return 0, fmt.Errorf("Not enough characters left for `%s` at %d, total size = %d", target, p.curPos, p.dataLen)
	}

	if p.data[p.curPos:p.curPos+l] == target {
		return l, nil
	}

	return 0, fmt.Errorf("No `%s` starting at %d", target, p.curPos)
}

func (p *Parser) ParseMulExpr() (int, int, int, error) {
	consumed, err := p.ParseText("mul")

	start := p.curPos

	if err != nil {
		log.Println(err)
		p.curPos += 1
		return 0, 0, start, err
	} else {
		log.Printf("Matched `mul` at %d:%d\n", p.curPos, p.curPos+consumed)
		p.curPos += consumed
	}

	err = p.ParseParen(false)
	p.curPos += 1

	if err != nil {
		log.Println(err)
		return 0, 0, start, err
	} else {
		log.Printf("Matched `(` at %d\n", p.curPos-1)
	}

	left, consumed, err := p.ParseInt()

	if err != nil {
		log.Println(err)
		p.curPos += 1
		return 0, 0, start, err
	} else {
		log.Printf("Matched %d at %d:%d\n", left, p.curPos, p.curPos+consumed)
		p.curPos += consumed
	}

	err = p.ParseComma()
	p.curPos += 1

	if err != nil {
		log.Println(err)
		return 0, 0, start, err
	} else {
		log.Printf("Matched , at %d\n", p.curPos-1)
	}

	right, consumed, err := p.ParseInt()

	if err != nil {
		log.Println(err)
		p.curPos += 1
		return 0, 0, start, err
	} else {
		log.Printf("Matched %d at %d:%d\n", right, p.curPos, p.curPos+consumed)
		p.curPos += consumed
	}

	err = p.ParseParen(true)
	p.curPos += 1

	if err != nil {
		log.Println(err)
		return 0, 0, start, err
	} else {
		log.Printf("Matched `)` at %d\n", p.curPos-1)
	}

	log.Printf("Matched expression `mul(%d,%d)`", left, right)
	return left, right, start, nil
}

func main() {
	content := parseInput()

	parser := Parser{data: content, dataLen: len(content), curPos: 0}

	result := 0

	matches := 0

	enabled := true

	for parser.curPos < parser.dataLen-1 {
		if enabled == true {
			left, right, start, mulExpErr := parser.ParseMulExpr()

			if mulExpErr == nil {
				matches += 1
				result += left * right
				continue
			}

			parser.curPos = start
		}

		consumed, err := parser.ParseText("do()")

		if err == nil {
			enabled = true
			parser.curPos += consumed
			continue
		}

		consumed, err = parser.ParseText("don't()")

		if err == nil {
			enabled = false
			parser.curPos += consumed
			continue
		}

		parser.curPos += 1
	}

	fmt.Printf("Result = %d, matches = %d\n", result, matches)
}
