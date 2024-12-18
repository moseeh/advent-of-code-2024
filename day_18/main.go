package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	y, x int
}

type QueueItem struct {
	point    Point
	distance int
	parent   *QueueItem
}

var (
	isCorrupted = make(map[Point]bool)
	startPoint  = Point{0, 0}
	endPoint    = Point{70, 70}
	gridSize    = 71
)

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	processInput(string(filebytes))

	// Find shortest path using BFS
	if path := findShortestPath(); len(path) > 0 {
		fmt.Printf("Shortest path found with length %d: %v\n", len(path)-1, path)
	} else {
		fmt.Println("No valid path found")
	}
}

func processInput(filestr string) {
	filearr := strings.Split(filestr, "\n")
	for i, line := range filearr {
		if i == 1024 {
			break
		}
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		isCorrupted[Point{y, x}] = true
	}
}

func findShortestPath() []Point {
	visited := make(map[Point]bool)
	queue := []*QueueItem{{point: startPoint, distance: 0, parent: nil}}
	visited[startPoint] = true

	directions := []Point{
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
		{-1, 0}, // up
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.point == endPoint {
			return reconstructPath(current)
		}

		for _, d := range directions {
			next := Point{
				y: current.point.y + d.y,
				x: current.point.x + d.x,
			}
			if isValidPoint(next) && !visited[next] {
				visited[next] = true
				queue = append(queue, &QueueItem{
					point:    next,
					distance: current.distance + 1,
					parent:   current,
				})
			}
		}
	}

	return []Point{}
}

func isValidPoint(p Point) bool {
	if p.x < 0 || p.x >= gridSize || p.y < 0 || p.y >= gridSize {
		return false
	}

	if isCorrupted[p] {
		return false
	}

	return true
}

func reconstructPath(endItem *QueueItem) []Point {
	path := []Point{}
	current := endItem

	for current != nil {
		path = append([]Point{current.point}, path...)
		current = current.parent
	}

	return path
}
