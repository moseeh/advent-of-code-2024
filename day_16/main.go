package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	y, x int
}

type QueueItem struct {
	point     Point
	turns     int
	direction string
	parent    *QueueItem
	index     int // for heap implementation
}

// Priority Queue implementation
type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].turns < pq[j].turns
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

var (
	isWall     = make(map[Point]bool)
	startPoint = Point{}
	endPoint   = Point{}
	gridSize   int
)

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filestr := string(filebytes)
	filearr := strings.Split(filestr, "\n")
	gridSize = len(filearr)
	processInput(filearr)

	path := findPathWithLeastTurns()
	fmt.Println(path)
	fmt.Println(calculateScore(path))
}

func calculateScore(path []Point) int {
	currentDirection := "right"
	score := 0

	getDirection := func(curr, next Point) string {
		if curr.y == next.y {
			if curr.x < next.x {
				return "right"
			}
			return "left"
		}
		if curr.x == next.x {
			if curr.y < next.y {
				return "down"
			}
			return "up"
		}
		return "none"
	}

	for i := 0; i < len(path)-1; i++ {
		score += 1
		direction := getDirection(path[i], path[i+1])
		if direction != currentDirection {
			score += 1000
			currentDirection = direction
		}
	}
	return score
}

func processInput(input []string) {
	for i, row := range input {
		for j, c := range row {
			if c == '#' {
				isWall[Point{i, j}] = true
			}
			if c == 'S' {
				startPoint = Point{i, j}
			}
			if c == 'E' {
				endPoint = Point{i, j}
			}
		}
	}
}

func findPathWithLeastTurns() []Point {
	// Initialize priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Initial state with no direction (first move doesn't count as turn)
	heap.Push(&pq, &QueueItem{
		point:     startPoint,
		turns:     0,
		direction: "",
		parent:    nil,
	})

	// Track visited states with turns
	visited := make(map[string]int) // key: "y,x,direction", value: minimum turns to reach this state
	visited[fmt.Sprintf("%d,%d,%s", startPoint.y, startPoint.x, "")] = 0

	// Define possible directions
	directions := []struct {
		delta Point
		name  string
	}{
		{Point{0, 1}, "right"},
		{Point{1, 0}, "down"},
		{Point{0, -1}, "left"},
		{Point{-1, 0}, "up"},
	}

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*QueueItem)

		// If we reached the end, this is the path with least turns (due to priority queue)
		if current.point == endPoint {
			return reconstructPath(current)
		}

		for _, d := range directions {
			next := Point{
				y: current.point.y + d.delta.y,
				x: current.point.x + d.delta.x,
			}

			if !isValidPoint(next) {
				continue
			}

			// Calculate new turns
			newTurns := current.turns
			if current.direction != "" && current.direction != d.name {
				newTurns++
			}

			// Create state key
			stateKey := fmt.Sprintf("%d,%d,%s", next.y, next.x, d.name)

			// Skip if we've found this state with fewer or equal turns
			if prevTurns, exists := visited[stateKey]; exists && prevTurns <= newTurns {
				continue
			}

			// Update visited and add to queue
			visited[stateKey] = newTurns
			heap.Push(&pq, &QueueItem{
				point:     next,
				turns:     newTurns,
				direction: d.name,
				parent:    current,
			})
		}
	}

	return []Point{}
}

func isValidPoint(p Point) bool {
	if p.x < 0 || p.x >= gridSize || p.y < 0 || p.y >= gridSize {
		return false
	}
	if isWall[p] {
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
