package main

import (
	"fmt"
	"os"
	"strings"
)

var Count int

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filestr := string(filebytes)
	filearr := strings.Split(filestr, "\n")
	Part1(filearr)
	fmt.Println(Count)
}

func Part1(arr []string) {
	for i, line := range arr {
		for j, c := range line {
			if c == '0' {
				visited := make(map[[2]int]bool)
				Part1helper(i, j, arr, c, visited)
			}
		}
	}
}

func Part1helper(i, j int, arr []string, c rune, visited map[[2]int]bool) {
	visited[[2]int{i, j}] = true

	if c == '9' {
		Count++
		return
	}

	directions := []struct{ dx, dy int }{
		{1, 0},
		{-1, 0},
		{0, -1},
		{0, 1},
	}

	for _, d := range directions {
		newI, newJ := i+d.dx, j+d.dy
		if newI >= 0 && newI < len(arr) && newJ >= 0 && newJ < len(arr[0]) && !visited[[2]int{newI, newJ}] && rune(arr[newI][newJ]) == c+1 {
			Part1helper(newI, newJ, arr, c+1, visited)
		}
	}
}
