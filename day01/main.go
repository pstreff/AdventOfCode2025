package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"log"
	"strconv"

	"github.com/pstreff/AdventOfCode2025/utils"
)

//go:embed *.txt
var txtFiles embed.FS

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

	result := utils.UInt99{Value: 50}
	password := 0

	for {
		line, ok := readLine()
		if !ok {
			break
		}

		direction := line[:1]
		number, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatalf("Failed to parse number: %s", err)
		}

		if direction == "L" {
			result.Sub(number)
		} else if direction == "R" {
			result.Add(number)
		}

		if int(result.Value) == 0 {
			password++
		}
	}

	fmt.Println("Part1")
	fmt.Printf("Result: %d\n", result.Value)
	fmt.Printf("Password: %d\n", password)
	fmt.Println("---------------")
}

func part2() {
	fileName := "input.txt"
	//fileName = "test.txt"

	readLine, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	result := utils.UInt99{Value: 50}

	for {
		line, ok := readLine()
		if !ok {
			break
		}

		direction := line[:1]
		number, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatalf("Failed to parse number: %s", err)
		}

		if direction == "L" {
			result.Sub(number)
		} else if direction == "R" {
			result.Add(number)
		}
	}

	fmt.Println("Part2")
	fmt.Printf("Result: %d\n", result.Value)
	fmt.Printf("Password: %d\n", result.Overflow)
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
