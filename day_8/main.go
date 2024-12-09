package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filestr := string(filebytes)
	filearr := strings.Split(filestr, "\n")
	antmap := make(map[[2]int]rune)

	for y, line := range filearr {
		for x, c := range line {
			if c != '.' {
				antmap[[2]int{y, x}] = c
			}
		}
	}

	fmt.Println(Part1(antmap, filearr))
	fmt.Println(Part2(antmap, filearr))
}

func Part1(antmap map[[2]int]rune, arr []string) int {
	locations := make(map[[2]int]int)
	for y, line := range arr {
		for x, c := range line {
			if c == '.' {
				continue
			}
			for k, v := range antmap {
				if v == c {

					dy, dx := y-k[0], x-k[1]
					if dx == 0 && dy == 0 {
						continue
					}
					dy, dx = y+dy, x+dx
					if (dy <= len(arr)-1 && dy >= 0) && (dx <= len(line)-1 && dx >= 0) {
						locations[[2]int{dy, dx}]++
					}
				}
			}
		}
	}
	return len(locations)
}

func Part2(antmap map[[2]int]rune, arr []string) int {
	locations := make(map[[2]int]int)
	for y, line := range arr {
		for x, c := range line {
			if c == '.' {
				continue
			}
			for k, v := range antmap {
				if v == c {
					dy, dx := y-k[0], x-k[1]
					if dx == 0 && dy == 0 {
						continue
					}
					newy, newx := y, x
					for {
						if !((newy <= len(arr)-1 && newy >= 0) && (newx <= len(line)-1 && newx >= 0)) {
							break
						}

						locations[[2]int{newy, newx}]++
						newy, newx = newy+dy, newx+dx
					}
				}
			}
		}
	}
	return len(locations)
}
