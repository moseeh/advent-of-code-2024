package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Height = 103
	Width  = 101
)

var (
	Coordinates    = make(map[string][2]int)
	Velocities     = make(map[string][2]int)
	FinalPositions = make(map[[2]int]int)
)

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filestr := string(filebytes)
	filearr := strings.Split(filestr, "\n")

	for i, line := range filearr {
		parts := strings.Split(line, " ")
		Coordinates[fmt.Sprintf("Robot num%d", i+1)] = processRobotinfo(parts[0])
		Velocities[fmt.Sprintf("Robot num%d", i+1)] = processRobotinfo(parts[1])
	}
	Motions()
	fmt.Println(Part1())
}

func Part1() int {
	for _, v := range Coordinates {
		FinalPositions[v]++
	}
	xlow, ylow := Width/2-1, Height/2-1
	quad1, quad2, quad3, quad4 := 0, 0, 0, 0

	for j := 0; j <= xlow; j++ {
		for i := 0; i <= ylow; i++ {
			quad1 += FinalPositions[[2]int{j, i}]
		}
	}
	for j := xlow + 2; j < Width; j++ {
		for i := 0; i <= ylow; i++ {
			quad2 += FinalPositions[[2]int{j, i}]
		}
	}
	for j := 0; j <= xlow; j++ {
		for i := ylow + 2; i < Height; i++ {
			quad3 += FinalPositions[[2]int{j, i}]
		}
	}
	for j := xlow + 2; j < Width; j++ {
		for i := ylow + 2; i < Height; i++ {
			quad4 += FinalPositions[[2]int{j, i}]
		}
	}
	fmt.Println(quad1, quad2, quad3, quad4)

	return quad1 * quad2 * quad3 * quad4
}

func Motions() {
	count := 0

	for count < 100 {
		for robot, coords := range Coordinates {
			x, y := coords[0]+Velocities[robot][0], coords[1]+Velocities[robot][1]
			if x < 0 {
				x = Width + x
			}
			if y < 0 {
				y = Height + y
			}
			if x >= Width {
				x = x - Width
			}
			if y >= Height {
				y = y - Height
			}
			Coordinates[robot] = [2]int{x, y}
		}
		count++
	}
}

func processRobotinfo(part string) [2]int {
	var arr [2]int
	coords := strings.Split(part[2:], ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	arr[0] = x
	arr[1] = y
	return arr
}
