package main

import (
	"fmt"
	"math"
)

func Part2(region [][2]int, arr []string) int {
	fmt.Println(region)
	if len(region) == 1 {
		return 4
	}
	// checked := make(map[[2]int]bool)
	first := region[0]
	last := region[len(region)-1]
	if first[0] == last[0] && first[1] == last[1]-1 && isPerfectSquare(len(region)) {
		return 4
	}
	directions := []struct{ dx, dy int }{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	var sides int = 1
	var currentDirection string
	getDirection := func(curr, next [2]int) string {
		if curr[0] == next[0] {
			if curr[1] < next[1] {
				return "right"
			}
			return "left"
		}
		if curr[1] == next[1] {
			if curr[0] < next[0] {
				return "down"
			}
			return "up"
		}
		return "none"
	}

	for i := 0; i < len(region)-1; i++ {
		curr := region[i]
		next := region[i+1]

		direction := getDirection(curr, next)
		I, J := curr[0], curr[1]

		if i == 0 {
			currentDirection = direction
			continue
		}

		if direction != currentDirection {
			for _, d := range directions {
				newI, newJ := I+d.dx, J+d.dy
				if newI < 0 || newI > len(arr)-1 {
					sides++
				}
				if newJ < 0 || newJ > len(arr[0])-1 {
					sides++
				}
				if newI >= 0 && newI < len(arr) && newJ >= 0 && newJ < len(arr[0]) {
					if arr[newI][newJ] != arr[curr[0]][curr[1]] {
						sides++
					}
				}
			}
			currentDirection = direction
		}
		I, J = next[0], next[1]
		if i == len(region)-2 {
			for _, d := range directions {
				newI, newJ := I+d.dx, J+d.dy
				if newI < 0 || newI > len(arr)-1 {
					sides++
				}
				if newJ < 0 || newJ > len(arr[0])-1 {
					sides++
				}
				if newI >= 0 && newI < len(arr) && newJ >= 0 && newJ < len(arr[0]) {
					if arr[newI][newJ] != arr[curr[0]][curr[1]] {
						sides++
					}
				}
			}
		}
	}
	fmt.Println(sides)
	if sides > 12 {
		sides -= 4
	}
	return sides
}

func isPerfectSquare(n int) bool {
	root := math.Sqrt(float64(n))

	return int(root)*int(root) == n
}
