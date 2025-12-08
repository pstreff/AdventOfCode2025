package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"log"
	"sort"
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

	readLine, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var result int64 = 0

	var freshIngredients [][]int

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		if in == "" {
			break
		}

		ingRange := strings.Split(in, "-")

		low, err := strconv.Atoi(ingRange[0])

		if err != nil {
			log.Fatal(err)
		}

		high, err := strconv.Atoi(ingRange[1])
		if err != nil {
			log.Fatal(err)
		}

		freshIngredients = append(freshIngredients, []int{low, high})
	}

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		check, err := strconv.Atoi(in)

		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(freshIngredients); i++ {
			if rangeContains(freshIngredients[i], check) {
				result++
				break
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

	readLine, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var result int64 = 0
	var freshIngredients [][]int

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		if in == "" {
			break
		}

		ingRange := strings.Split(in, "-")

		low, err := strconv.Atoi(ingRange[0])

		if err != nil {
			log.Fatal(err)
		}

		high, err := strconv.Atoi(ingRange[1])
		if err != nil {
			log.Fatal(err)
		}

		freshIngredients = append(freshIngredients, []int{low, high})
	}

	sort.Slice(freshIngredients, func(i, j int) bool {
		if freshIngredients[i][0] != freshIngredients[j][0] {
			return freshIngredients[i][0] < freshIngredients[j][0]
		}
		return freshIngredients[i][1] < freshIngredients[j][1]
	})

	freshIngredients = mergeRanges(freshIngredients)

	for i := 0; i < len(freshIngredients); i++ {
		result += int64((freshIngredients[i][1] - freshIngredients[i][0]) + 1)
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

func rangeContains(r []int, num int) bool {
	low := r[0]
	high := r[1]

	if num >= low && num <= high {
		return true
	}
	return false
}

func mergeRanges(l [][]int) [][]int {
	if len(l) == 0 {
		return l
	}

	merged := [][]int{l[0]}

	for i := 1; i < len(l); i++ {
		last := merged[len(merged)-1]
		current := l[i]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}
