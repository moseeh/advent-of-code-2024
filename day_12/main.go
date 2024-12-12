package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	PartRegion = make(map[[2]int]string)
	Regions    = make(map[string][][2]int)
	RegionNo   = 0
	Perimeter  = make(map[string]int)
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
	sum := 0
	for k, v := range Regions {
		fmt.Printf("The area and perimeter of %s, is %d and %d respectively\n", k, len(v), Perimeter[k])
		sum += len(v)*Perimeter[k]
	}
	fmt.Println(sum)
}

func Part1(arr []string) {
	for i, line := range arr {
		for j, c := range line {
			if _, exists := PartRegion[[2]int{i, j}]; exists {
				continue
			}
			RegionNo++
			Regions[fmt.Sprintf("Region %d", RegionNo)] = append(Regions[fmt.Sprintf("Region %d", RegionNo)], [2]int{i, j})
			PartRegion[[2]int{i, j}] = fmt.Sprintf("Region %d", RegionNo)
			Part1helper(arr, i, j, c)
		}
	}
}

func Part1helper(arr []string, i, j int, c rune) {
	// hasadjacent := false
	directions := []struct{ dx, dy int }{
		{1, 0},
		{-1, 0},
		{0, -1},
		{0, 1},
	}

	for _, d := range directions {
		newI, newJ := i+d.dx, j+d.dy
		if newI >= 0 && newI < len(arr) && newJ >= 0 && newJ < len(arr[0]) {
			if rune(arr[newI][newJ]) == c {
				// hasadjacent = true
				if _, exists := PartRegion[[2]int{newI, newJ}]; !exists {
					PartRegion[[2]int{newI, newJ}] = fmt.Sprintf("Region %d", RegionNo)
					Regions[fmt.Sprintf("Region %d", RegionNo)] = append(Regions[fmt.Sprintf("Region %d", RegionNo)], [2]int{newI, newJ})
					Part1helper(arr, newI, newJ, c)

				}
			} else {
				Perimeter[fmt.Sprintf("Region %d", RegionNo)]++
			}
		}
	}
	if i == 0 || i == len(arr)-1 {
		Perimeter[fmt.Sprintf("Region %d", RegionNo)]++
	}
	if j == 0 || j == len(arr[0])-1 {
		Perimeter[fmt.Sprintf("Region %d", RegionNo)]++
	}
}
