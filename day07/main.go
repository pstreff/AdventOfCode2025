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

		cells := strings.Split(string(in), "")
		grid = append(grid, cells)
	}

	startPos := -1
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == "S" {
			startPos = i
		}
	}

	if startPos == -1 {
		panic("No starting position")
	}

	grid[1][startPos] = "|"

	for i := 2; i < len(grid); i++ {
		hasSplitter := hasSplitter(grid[i])
		for j := 0; j < len(grid[i]); j++ {
			if hasSplitter {
				if grid[i][j] == "^" && grid[i-1][j] == "|" {
					if grid[i][j-1] != "^" {
						grid[i][j-1] = "|"
					}
					if grid[i][j+1] != "^" {
						grid[i][j+1] = "|"
					}
					result++
				}
				if grid[i-1][j] == "|" && grid[i][j] == "." {
					grid[i][j] = "|"
				}
			} else {
				if grid[i-1][j] == "|" {
					grid[i][j] = "|"
				}
			}
		}
	}

	//for i := 0; i < len(grid); i++ {
	//	fmt.Println(grid[i])
	//}

	fmt.Println("Part1")
	fmt.Printf("Result: %d\n", result)
	fmt.Println("---------------")
}

type Beam struct {
	Ways  int
	Point Point
}

type Point struct {
	X, Y int
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
	var beams [][]Beam

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		cells := strings.Split(string(in), "")
		grid = append(grid, cells)

		var bs []Beam
		for range cells {
			bs = append(bs, Beam{Ways: 0})
		}
		beams = append(beams, bs)
	}

	startPos := -1
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == "S" {
			startPos = i
		}
	}

	if startPos == -1 {
		panic("No starting position")
	}

	grid[1][startPos] = "|"
	beams[1][startPos].Ways = 1

	for i := 2; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "^" && grid[i-1][j] == "|" {
				previous := beams[i-1][j]

				if grid[i][j-1] != "^" {
					grid[i][j-1] = "|"

					if previous.Ways == 0 {
						beams[i][j-1].Ways += 1
					} else {
						beams[i][j-1].Ways += previous.Ways
					}
				}
				if grid[i][j+1] != "^" {
					grid[i][j+1] = "|"

					if previous.Ways == 0 {
						beams[i][j+1].Ways += 1
					} else {
						beams[i][j+1].Ways += previous.Ways
					}
				}
			}
			if grid[i-1][j] == "|" && grid[i][j] != "^" {
				grid[i][j] = "|"
				previous := beams[i-1][j]
				beams[i][j].Ways += previous.Ways
			}
		}
	}

	for i := 0; i < len(grid[len(grid)-1]); i++ {
		if grid[len(grid)-1][i] == "|" {
			b := beams[len(grid)-1][i]
			result += int64(b.Ways)
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

func hasSplitter(l []string) bool {
	for i := 0; i < len(l); i++ {
		if l[i] == "^" {
			return true
		}
	}
	return false
}
