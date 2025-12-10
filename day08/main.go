package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed *.txt
var txtFiles embed.FS

func main() {
	//part1()
	part2()
}

type Junction struct {
	X, Y, Z     int
	Connections []*Junction
}

func (j *Junction) Equals(other *Junction) bool {
	if j == nil || other == nil {
		return j == other
	}

	return j.X == other.X && j.Y == other.Y && j.Z == other.Z
}

func (j *Junction) HasDirectConnectionTo(other *Junction) bool {
	for _, conn := range j.Connections {
		if conn.Equals(other) {
			return true
		}
	}
	for _, conn := range other.Connections {
		if conn.Equals(j) {
			return true
		}
	}
	return false
}

func part1() {
	fileName := "input.txt"
	//fileName = "test.txt"

	readLine, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var result int64 = 1

	var junctions []Junction

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		coords := strings.Split(string(in), ",")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatalf("Failed to parse coordinate: %s", err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatalf("Failed to parse coordinate: %s", err)
		}
		z, err := strconv.Atoi(coords[2])
		if err != nil {
			log.Fatalf("Failed to parse coordinate: %s", err)
		}

		junctions = append(junctions, Junction{X: x, Y: y, Z: z})
	}

	circuits := make(map[int][]*Junction)
	mapNextIndex := 0

	for _, j := range junctions {
		circuits[mapNextIndex] = append(circuits[mapNextIndex], &j)
		mapNextIndex += 1
	}

	var a, b *Junction

	for c := 0; c < 1000; c++ {
		d := math.MaxFloat64
		for i := 0; i < len(junctions); i++ {
			candidatA := &junctions[i]
			for j := 0; j < len(junctions); j++ {
				candidatB := &junctions[j]
				if i == j {
					continue
				}

				if candidatA.HasDirectConnectionTo(candidatB) {
					continue
				}

				distance := getDistance(junctions[i], junctions[j])
				if distance < d {
					d = distance
					a = candidatA
					b = candidatB
				}
			}
		}

		if a == nil || b == nil {
			panic("Failed to find a and b")
		}

		a.Connections = append(a.Connections, b)

		for i := 0; i < mapNextIndex; i++ {
			if circuits[i] == nil {
				continue
			}

			if junctionInList(a, circuits[i]) {
				for m := 0; m < mapNextIndex; m++ {
					if m == i || circuits[m] == nil {
						continue
					}

					if junctionInList(b, circuits[m]) {
						circuits[i] = append(circuits[i], circuits[m]...)
						delete(circuits, m)
						break
					}
				}

				if !junctionInList(b, circuits[i]) {
					circuits[i] = append(circuits[i], b)
				}
			} else if junctionInList(b, circuits[i]) {
				for m := 0; m < mapNextIndex; m++ {
					if m == i || circuits[m] == nil {
						continue
					}

					if junctionInList(a, circuits[m]) {
						circuits[i] = append(circuits[i], circuits[m]...)
						delete(circuits, m)
						break
					}
				}
				if !junctionInList(a, circuits[i]) {
					circuits[i] = append(circuits[i], a)
				}
			}
		}
	}

	valuesLengths := make([]int, len(circuits))
	for v := 0; v < len(circuits); v++ {
		valuesLengths = append(valuesLengths, len(circuits[v]))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(valuesLengths)))

	fmt.Println(valuesLengths)

	top3 := valuesLengths[:3]

	for i := 0; i < len(top3); i++ {
		fmt.Println(top3[i])
		result *= int64(top3[i])
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

	var result int64 = 1

	var junctions []Junction

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		coords := strings.Split(string(in), ",")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatalf("Failed to parse coordinate: %s", err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatalf("Failed to parse coordinate: %s", err)
		}
		z, err := strconv.Atoi(coords[2])
		if err != nil {
			log.Fatalf("Failed to parse coordinate: %s", err)
		}

		junctions = append(junctions, Junction{X: x, Y: y, Z: z})
	}

	circuits := make(map[int][]*Junction)
	mapNextIndex := 0

	for _, j := range junctions {
		circuits[mapNextIndex] = append(circuits[mapNextIndex], &j)
		mapNextIndex += 1
	}

	var a, b *Junction

	for {
		d := math.MaxFloat64
		for i := 0; i < len(junctions); i++ {
			candidatA := &junctions[i]
			for j := 0; j < len(junctions); j++ {
				candidatB := &junctions[j]
				if i == j {
					continue
				}

				if candidatA.HasDirectConnectionTo(candidatB) {
					continue
				}

				distance := getDistance(junctions[i], junctions[j])
				if distance < d {
					d = distance
					a = candidatA
					b = candidatB
				}
			}
		}

		if a == nil || b == nil {
			panic("Failed to find a and b")
		}

		a.Connections = append(a.Connections, b)

		for i := 0; i < mapNextIndex; i++ {
			if circuits[i] == nil {
				continue
			}

			if junctionInList(a, circuits[i]) {
				for m := 0; m < mapNextIndex; m++ {
					if m == i || circuits[m] == nil {
						continue
					}

					if junctionInList(b, circuits[m]) {
						circuits[i] = append(circuits[i], circuits[m]...)
						delete(circuits, m)
						break
					}
				}

				if !junctionInList(b, circuits[i]) {
					circuits[i] = append(circuits[i], b)
				}
			} else if junctionInList(b, circuits[i]) {
				for m := 0; m < mapNextIndex; m++ {
					if m == i || circuits[m] == nil {
						continue
					}

					if junctionInList(a, circuits[m]) {
						circuits[i] = append(circuits[i], circuits[m]...)
						delete(circuits, m)
						break
					}
				}
				if !junctionInList(a, circuits[i]) {
					circuits[i] = append(circuits[i], a)
				}
			}
		}

		if len(circuits) == 1 {
			fmt.Println(a, b)
			result = int64(a.X * b.X)
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

func getDistance(a Junction, b Junction) float64 {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	dz := float64(a.Z - b.Z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func junctionInList(j *Junction, l []*Junction) bool {
	for i := 0; i < len(l); i++ {
		if l[i].Equals(j) {
			return true
		}
	}
	return false
}
