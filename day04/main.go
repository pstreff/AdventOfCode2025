package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"log"
	"strings"
)

//go:embed *.txt
var txtFiles embed.FS

type Coord struct {
	x int
	y int
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

	var grid [][]string

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		var l []string
		l = strings.Split(in, "")

		grid = append(grid, l)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "@" {
				n := checkNeighbours(i, j, grid)
				if n < 4 {
					result += 1
				}
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

	var grid [][]string

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		var l []string
		l = strings.Split(in, "")

		grid = append(grid, l)
	}

	for {
		var toRemove []Coord

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == "@" {
					n := checkNeighbours(i, j, grid)
					if n < 4 {
						result += 1
						toRemove = append(toRemove, Coord{x: i, y: j})
					}
				}
			}
		}

		if len(toRemove) > 0 {
			for i := 0; i < len(toRemove); i++ {
				grid[toRemove[i].x][toRemove[i].y] = "."
			}
		} else if len(toRemove) == 0 {
			break
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

	scanner := bufio.NewScanner(bytes.NewReader(data))

	next := func() (string, bool) {
		if scanner.Scan() {
			return scanner.Text(), true
		}
		return "", false
	}

	return next, nil
}

func checkNeighbours(x, y int, grid [][]string) int {
	neighbours := 0

	if x-1 >= 0 && y-1 >= 0 {
		if grid[x-1][y-1] == "@" {
			neighbours++
		}
	}

	if x-1 >= 0 {
		if grid[x-1][y] == "@" {
			neighbours++
		}
	}

	if x-1 >= 0 && y+1 < len(grid[x-1]) {
		if grid[x-1][y+1] == "@" {
			neighbours++
		}
	}

	if y-1 >= 0 {
		if grid[x][y-1] == "@" {
			neighbours++
		}
	}

	if y+1 < len(grid[x]) {
		if grid[x][y+1] == "@" {
			neighbours++
		}
	}

	if y-1 >= 0 && x+1 < len(grid) {
		if grid[x+1][y-1] == "@" {
			neighbours++
		}
	}

	if x+1 < len(grid) {
		if grid[x+1][y] == "@" {
			neighbours++
		}
	}

	if x+1 < len(grid) && y+1 < len(grid[x+1]) {
		if grid[x+1][y+1] == "@" {
			neighbours++
		}
	}

	return neighbours
}
