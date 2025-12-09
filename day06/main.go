package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed *.txt
var txtFiles embed.FS

type Equation struct {
	Terms  []int
	Op     string
	Result int64
}

func main() {
	part1()
	part2()
}

func part1() {
	fileName := "input.txt"
	//fileName = "test.txt"

	readLine, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var result int64 = 0

	firstLine, ok := readLine()
	if !ok {
		panic("Failed to read first line")
	}

	terms := strings.Fields(firstLine)
	equations := make([]Equation, len(terms))

	for i := 0; i < len(terms); i++ {
		num, err := strconv.Atoi(terms[i])
		if err != nil {
			panic(err)
		}

		equations[i].Terms = append(equations[i].Terms, num)
	}

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		terms = strings.Fields(in)
		_, err := strconv.Atoi(terms[0])
		if err != nil {
			for i := 0; i < len(terms); i++ {
				equations[i].Op = terms[i]
			}
		} else {
			for i := 0; i < len(terms); i++ {
				num, err := strconv.Atoi(terms[i])
				if err != nil {
					panic(err)
				}

				equations[i].Terms = append(equations[i].Terms, num)
			}
		}
	}

	for i := 0; i < len(equations); i++ {
		equations[i].Result = solveEquation(equations[i])
		result += equations[i].Result
	}

	fmt.Println("Part1")
	fmt.Printf("Result: %d\n", result)
	fmt.Println("---------------")
}

func part2() {
	fileName := "input.txt"
	//fileName = "test.txt"

	readLine, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var result int64 = 0

	var lines []string

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		lines = append(lines, in)
	}

	operators := lines[len(lines)-1]
	lines = lines[:len(lines)-1]
	equations := make([]Equation, len(strings.Fields(lines[0])))

	ops := strings.Fields(operators)

	for i := len(ops) - 1; i >= 0; i-- {
		equations[i].Op = ops[i]
	}

	equationIndex := len(equations) - 1
	for i := len(lines[0]) - 1; i >= 0; i-- {
		numberOfTerms := len(lines)

		var term string

		for j := 0; j < numberOfTerms; j++ {
			term += string(lines[j][i])
		}

		term = strings.TrimSpace(term)

		if term != "" {
			intTerm, err := strconv.Atoi(term)
			if err != nil {
				panic(err)
			}
			equations[equationIndex].Terms = append(equations[equationIndex].Terms, intTerm)
		} else {
			equationIndex--
		}

	}
	for i := 0; i < len(equations); i++ {
		equations[i].Result = solveEquation(equations[i])
		result += equations[i].Result
	}

	fmt.Println("Part2")
	fmt.Printf("Result: %d\n", result)
	fmt.Println("---------------")
}

func fileReader(path string) (func() (string, bool), error) {
	data, err := txtFiles.ReadFile(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))

	next := func() (string, bool) {
		if scanner.Scan() {
			return scanner.Text(), true
		}
		return "", false
	}

	return next, nil
}

func solveEquation(e Equation) int64 {
	result := int64(0)
	if e.Op == "*" {
		result = 1
	}

	for _, term := range e.Terms {
		if e.Op == "*" {
			result *= int64(term)
		} else {
			result += int64(term)
		}
	}
	return result
}
