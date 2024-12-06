package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	Count       = make(map[[2]int]int)
	Solved      bool
	Positions   int
	ObstacleMet = make(map[[2]int]bool)
	Up          = make(map[[2]int]bool)
	Down        = make(map[[2]int]bool)
	Right       = make(map[[2]int]bool)
	Left        = make(map[[2]int]bool)
)

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filestr := string(filebytes)
	filearr := strings.Split(filestr, "\n")
	Part1(filearr)
	Part2(filearr)
	fmt.Println(Positions)
}

func Part1(area []string) {
	for i, line := range area {
		for j, char := range line {
			if char == '^' {
				Count[[2]int{i, j}]++
				// fmt.Printf("guard on line %d row %d\n", i+1, j+1)
				tracepathup(area, i, j)
				break

			}
		}
	}
}

func tracepathup(area []string, row, j int) {
	for i := row - 1; i >= 0; i-- {
		if area[i][j] != '#' {
			Count[[2]int{i, j}]++
			if i == 0 {
				// fmt.Println("end")
				Solved = true
				break
			}
		} else {
			// fmt.Printf("obstacle on line %d row %d while going up\n", i+1, j+1)

			if ObstacleMet[[2]int{i, j}] && Up[[2]int{i, j}] {
				fmt.Println("is in a loop")
				break
			}
			ObstacleMet[[2]int{i, j}] = true
			Up[[2]int{i, j}] = true
			tracepathright(area[i+1], i+1, j, area)
			break
		}
	}
}

func tracepathright(line string, i, column int, area []string) {
	for j := column + 1; j <= len(line)-1; j++ {
		if line[j] != '#' {
			Count[[2]int{i, j}]++
			if j == len(line)-1 {
				Solved = true
				// fmt.Println("end")
				break
			}
		} else {
			// fmt.Printf("obstacle on line %d row %d while going right\n", i+1, j+1)
			if ObstacleMet[[2]int{i, j}] && Right[[2]int{i, j}] {
				fmt.Println("is in a loop")
				break
			}
			Right[[2]int{i, j}] = true
			ObstacleMet[[2]int{i, j}] = true
			tracepathdown(area, i+1, j-1)
			break
		}
	}
}

func tracepathdown(area []string, row, j int) {
	for i := row; i <= len(area)-1; i++ {
		if area[i][j] != '#' {
			Count[[2]int{i, j}]++
			if i == len(area)-1 {
				Solved = true
				// fmt.Println("end")
				break
			}
		} else {
			// fmt.Printf("obstacle on line %d row %d while going down\n", i+1, j+1)
			if ObstacleMet[[2]int{i, j}] && Down[[2]int{i, j}] {
				fmt.Println("is in a loop")
				break
			}
			ObstacleMet[[2]int{i, j}] = true
			Down[[2]int{i, j}] = true
			tracepathleft(area[i-1], i-1, j, area)
			break
		}
	}
}

func tracepathleft(line string, i, column int, area []string) {
	for j := column - 1; j >= 0; j-- {
		if line[j] != '#' {
			Count[[2]int{i, j}]++
			if j == 0 {
				Solved = true
				// fmt.Println("end")
				break
			}
		} else {
			// fmt.Printf("obstacle on line %d row %d while going left\n", i+1, j+1)
			if ObstacleMet[[2]int{i, j}] && Left[[2]int{i, j}] {
				fmt.Println("is in a loop")
				break
			}
			ObstacleMet[[2]int{i, j}] = true
			Left[[2]int{i, j}] = true
			tracepathup(area, i, j+1)
			break
		}
	}
}
