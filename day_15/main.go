package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	IsWall       = make(map[[2]int]bool)
	IsBox        = make(map[[2]int]bool)
	RobotPostion = [2]int{}
	Height       int
	Width        int
)

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filestr := string(filebytes)
	filearr := strings.Split(filestr, "\n\n")
	parseMap(filearr[0])
	Part1(filearr[1])
	sum := 0
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			if IsBox[[2]int{i, j}] {
				sum += 100*i + j
				fmt.Print("O")
			} else if IsWall[[2]int{i, j}] {
				fmt.Print("#")
			} else {
				if i == RobotPostion[0] && j == RobotPostion[1] {
					fmt.Print("@")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println(sum)
}

func parseMap(str string) {
	strarr := strings.Split(str, "\n")
	Height = len(strarr)
	Width = len(strarr[0])
	for i, row := range strarr {
		for j, char := range row {
			if char == '#' {
				IsWall[[2]int{i, j}] = true
			}
			if char == 'O' {
				IsBox[[2]int{i, j}] = true
			}
			if char == '@' {
				RobotPostion = [2]int{i, j}
			}
		}
	}
}

func Part1(directions string) {
	for _, c := range directions {
		i, j := RobotPostion[0], RobotPostion[1]
		fmt.Println(i, j)
		if c == '^' {
			for k := i - 1; k > 0; k-- {
				if !IsBox[[2]int{k, j}] && !IsWall[[2]int{k, j}] && k == i-1 {
					RobotPostion = [2]int{k, j}
					break
				}
				if !IsBox[[2]int{k, j}] && !IsWall[[2]int{k, j}] {
					for l := k; l < i; l++ {
						if l != i-1 {
							IsBox[[2]int{l, j}] = true
						} else {
							IsBox[[2]int{l, j}] = false
							RobotPostion = [2]int{l, j}
						}
					}
					break
				}
				if IsBox[[2]int{k, j}] {
					continue
				}
				if IsWall[[2]int{k, j}] {
					break
				}
			}
		}
		if c == 'v' {
			for k := i + 1; k < Height; k++ {
				if !IsBox[[2]int{k, j}] && !IsWall[[2]int{k, j}] && k == i+1 {
					RobotPostion = [2]int{k, j}
					break
				}
				if !IsBox[[2]int{k, j}] && !IsWall[[2]int{k, j}] {
					for l := k; l > i; l-- {
						if l != i+1 {
							IsBox[[2]int{l, j}] = true
						} else {
							IsBox[[2]int{l, j}] = false
							RobotPostion = [2]int{l, j}
						}
					}
					break
				}
				if IsBox[[2]int{k, j}] {
					continue
				}
				if IsWall[[2]int{k, j}] {
					break
				}
			}
		}
		if c == '<' {
			for k := j - 1; k > 0; k-- {
				if !IsBox[[2]int{i, k}] && !IsWall[[2]int{i, k}] && k == j-1 {
					RobotPostion = [2]int{i, k}
					break
				}
				if !IsBox[[2]int{i, k}] && !IsWall[[2]int{i, k}] {
					for l := k; l < j; l++ {
						if l != j-1 {
							IsBox[[2]int{i, l}] = true
						} else {
							IsBox[[2]int{i, l}] = false
							RobotPostion = [2]int{i, l}
						}
					}
					break
				}
				if IsBox[[2]int{i, k}] {
					continue
				}
				if IsWall[[2]int{i, k}] {
					break
				}
			}
		}
		if c == '>' {
			for k := j + 1; k < Width; k++ {
				if !IsBox[[2]int{i, k}] && !IsWall[[2]int{i, k}] && k == j+1 {
					RobotPostion = [2]int{i, k}
					break
				}
				if !IsBox[[2]int{i, k}] && !IsWall[[2]int{i, k}] {
					for l := k; l > j; l-- {
						if l != j+1 {
							IsBox[[2]int{i, l}] = true
						} else {
							IsBox[[2]int{i, l}] = false
							RobotPostion = [2]int{i, l}
						}
					}
					break
				}
				if IsBox[[2]int{i, k}] {
					continue
				}
				if IsWall[[2]int{i, k}] {
					break
				}
			}
		}
	}
}
