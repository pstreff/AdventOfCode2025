package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

//go:embed *.txt
var txtFiles embed.FS

func main() {
	part1()
	part2()
}

type Point struct {
	X, Y int
}

type Square struct {
	P1, P2 Point
}

//var pipCache = utils.NewLRU[Point, bool](5000000)

func (s Square) Area() int {
	return abs((abs(s.P2.X-s.P1.X) + 1) * (abs(s.P2.Y-s.P1.Y) + 1))
}

func part1() {
	fileName := "input.txt"
	//fileName = "test.txt"

	readLine, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	//var result int64 = 1

	var points []Point

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		coords := strings.Split(string(in), ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		points = append(points, Point{x, y})
	}

	largestArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if i == j {
				continue
			}

			sq := Square{points[i], points[j]}
			sqArea := sq.Area()

			//fmt.Printf("Square (%d,%d) (%d,%d) -> %d\n", points[i].X, points[i].Y, points[j].X, points[j].Y, sqArea)

			if sqArea > largestArea {
				largestArea = sqArea
			}
		}
	}

	fmt.Println("Part1")
	fmt.Printf("Result: %d\n", largestArea)
	fmt.Println("---------------")
}

func part2() {
	fileName := "input.txt"
	//fileName = "test.txt"

	readLine, err := fileReader(fileName)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var points []Point

	for {
		in, ok := readLine()
		if !ok {
			break
		}

		coords := strings.Split(string(in), ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		points = append(points, Point{x, y})
	}

	numWorkers := runtime.NumCPU()
	type job struct {
		i, j int
	}

	jobs := make(chan job, 1024)

	var largestArea int64 = 0
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for w := 0; w < numWorkers; w++ {
		go func() {
			defer wg.Done()
			for job := range jobs {
				sq := Square{points[job.i], points[job.j]}
				sqArea := sq.Area()

				currentMax := atomic.LoadInt64(&largestArea)
				if int64(sqArea) <= currentMax {
					continue
				}

				if !RectangleInsidePolygon(sq, points) {
					continue
				}

				for {
					currentMax := atomic.LoadInt64(&largestArea)
					if int64(sqArea) <= currentMax {
						break
					}
					if atomic.CompareAndSwapInt64(&largestArea, currentMax, int64(sqArea)) {
						break
					}
				}
			}
		}()
	}

	go func() {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				jobs <- job{i, j}
			}
		}
		close(jobs)
	}()

	wg.Wait()

	fmt.Println("Part2")
	fmt.Printf("Result: %d\n", largestArea)
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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func PointInPolygon(p Point, polygon []Point) bool {
	inside := false
	n := len(polygon)
	j := n - 1

	for i := 0; i < n; i++ {
		pi := polygon[i]
		pj := polygon[j]

		if pointOnSegment(p, pj, pi) {
			//pipCache.Put(p, true)
			return true
		}

		if (pi.Y > p.Y) != (pj.Y > p.Y) && p.X < (pj.X-pi.X)*(p.Y-pi.Y)/(pj.Y-pi.Y)+pi.X {
			inside = !inside
		}
		j = i
	}

	return inside
}

func pointOnSegment(p, a, b Point) bool {
	minX, maxX := min(a.X, b.X), max(a.X, b.X)
	minY, maxY := min(a.Y, b.Y), max(a.Y, b.Y)

	if p.X < minX || p.X > maxX || p.Y < minY || p.Y > maxY {
		return false
	}

	return (b.X-a.X)*(p.Y-a.Y) == (b.Y-a.Y)*(p.X-a.X)
}

func RectangleInsidePolygon(sq Square, polygon []Point) bool {
	minX, maxX := min(sq.P1.X, sq.P2.X), max(sq.P1.X, sq.P2.X)
	minY, maxY := min(sq.P1.Y, sq.P2.Y), max(sq.P1.Y, sq.P2.Y)

	for x := minX; x <= maxX; x++ {
		if !PointInPolygon(Point{x, minY}, polygon) {
			return false
		}

		if !PointInPolygon(Point{x, maxY}, polygon) {
			return false
		}
	}

	for y := minY + 1; y < maxY; y++ {
		if !PointInPolygon(Point{minX, y}, polygon) {
			return false
		}

		if !PointInPolygon(Point{minX, y}, polygon) {
			return false
		}
	}
	return true
}
