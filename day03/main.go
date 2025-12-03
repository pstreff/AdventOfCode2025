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

func main() {
	part1()
	part2()
}

func part1() {
	fileName := "input.txt"
	//fileName = "test.txt"

	digits := 2

	readLine, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var result int64 = 0

	for {
		bank, ok := readLine()
		if !ok {
			break
		}

		var joltageSlice []int
		batteries := strings.Split(string(bank), "")

		joltageSlice = findJoltage(batteries, digits)

		joltageStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(joltageSlice)), ""), "[]")

		joltage, err := strconv.Atoi(joltageStr)

		if err != nil {
			log.Fatal(err)
		}

		result += int64(joltage)
	}

	fmt.Println("Part1")
	fmt.Printf("Result: %d\n", result)
	fmt.Println("---------------")
}

func part2() {

	fileName := "input.txt"
	//fileName = "test.txt"

	digits := 12

	readLine, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var result int64 = 0

	for {
		bank, ok := readLine()
		if !ok {
			break
		}

		var joltageSlice []int
		batteries := strings.Split(string(bank), "")

		joltageSlice = findJoltage(batteries, digits)

		joltageStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(joltageSlice)), ""), "[]")

		joltage, err := strconv.Atoi(joltageStr)

		if err != nil {
			log.Fatal(err)
		}

		result += int64(joltage)
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

func findJoltage(bank []string, digits int) []int {
	if digits <= 0 || len(bank) < digits || len(bank) == 0 {
		return []int{}
	}

	largest := -1
	largestIndex := -1
	s := make([]int, 0, digits)

	for i := 0; i <= (len(bank) - digits); i++ {
		value, err := strconv.Atoi(bank[i])
		if err != nil {
			log.Fatal(err)
		}

		if value > largest {
			largest = value
			largestIndex = i
		}
	}

	var newBank []string
	if largestIndex+1 < len(bank) {
		newBank = bank[largestIndex+1:]
	}

	j := findJoltage(newBank, digits-1)
	s = append(s, largest)
	s = append(s, j...)
	return s
}
