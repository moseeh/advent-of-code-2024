package main

import (
	"fmt"
	"os"
	"strings"
)

var Count = make(map[[2]int]int)

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filestr := string(filebytes)
	filearr := strings.Split(filestr, "\n")
	Part1(filearr)
	fmt.Println(len(Count))
}

func Part1(area []string) {
	for i, line := range area {
		for j, char := range line {
			if char == '^' {
				Count[[2]int{i, j}]++
				fmt.Printf("guard on line %d row %d\n", i+1, j+1)
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
				fmt.Println("end")
				break
			}
		} else {
			fmt.Printf("obstacle on line %d row %d\n", i+1, j+1)
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
				fmt.Println("end")
				break
			}
		} else {
			fmt.Printf("obstacle on line %d row %d\n", i+1, j+1)
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
				fmt.Println("end")
				break
			}
		} else {
			fmt.Printf("obstacle on line %d row %d\n", i+1, j+1)

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
				fmt.Println("end")
				break
			}
		} else {
			fmt.Printf("obstacle on line %d row %d\n", i+1, j+1)
			tracepathup(area, i, j+1)
			break
		}
	}
}
