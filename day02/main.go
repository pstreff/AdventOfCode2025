package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"io"
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

	getNextRange, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var result int64 = 0

	for {
		idRange, ok := getNextRange()
		if !ok {
			break
		}

		parts := strings.Split(idRange, "-")

		if len(parts) != 2 {
			panic("Invalid id range format")
		}

		startStr := parts[0]
		endStr := parts[1]

		start, err := strconv.Atoi(startStr)
		if err != nil {
			panic("Invalid start number")
		}
		end, err := strconv.Atoi(endStr)
		if err != nil {
			panic("Invalid end number")
		}

		for i := int64(start); i <= int64(end); i++ {
			digits := digitCount(i)
			if digits%2 != 0 {
				continue
			}

			middle := digits / 2
			number := strconv.FormatInt(i, 10)
			left := number[:middle]
			right := number[middle:]

			if left == right {
				result += i
			}
		}
	}

	fmt.Println("Part1")
	fmt.Printf("Result: %d\n", result)
	fmt.Println("---------------")
}

func part2() {
	fileName := "input.txt"
	//fileName = "test.txt"

	getNextRange, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var result int64 = 0

	for {
		idRange, ok := getNextRange()
		if !ok {
			break
		}

		parts := strings.Split(idRange, "-")

		if len(parts) != 2 {
			panic("Invalid id range format")
		}

		startStr := parts[0]
		endStr := parts[1]

		start, err := strconv.Atoi(startStr)
		if err != nil {
			panic("Invalid start number")
		}
		end, err := strconv.Atoi(endStr)
		if err != nil {
			panic("Invalid end number")
		}

		for i := int64(start); i <= int64(end); i++ {
			digits := digitCount(i)
			invalid := false

			for size := 1; size <= digits; size++ {
				if invalid {
					continue
				}

				if digits%size != 0 {
					continue
				}

				var checks []string
				number := strconv.FormatInt(i, 10)
				for y := 0; y < digits; y += size {
					part := number[y : y+size]
					checks = append(checks, part)
				}

				if len(checks) > 1 && allSameStrings(checks) {
					result += i
					invalid = true
				}
			}
		}
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

	reader := bufio.NewReader(bytes.NewReader(data))

	next := func() (string, bool) {
		part, err := reader.ReadString(',')

		if err == io.EOF {
			if len(part) > 0 {
				return strings.TrimSpace(part), true
			}
			return "", false
		}

		if err != nil {
			return "", false
		}

		part = strings.TrimSuffix(part, ",")

		return strings.TrimSpace(part), true
	}

	return next, nil
}

func digitCount(num int64) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}

func allSameStrings(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}
