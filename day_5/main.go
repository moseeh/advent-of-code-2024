package main

import (
	"fmt"
	"os"
	"strconv"
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
	updates := [][]int{}
	ordering := make(map[int][]int)
	for _, str := range filearr {
		if str == "" {
			continue
		}
		if strings.Contains(str, "|") {
			order := strToINt(strings.Split(str, "|"))
			ordering[order[1]] = append(ordering[order[1]], order[0])
		}
		if strings.Contains(str, ",") {
			update := strToINt(strings.Split(str, ","))
			updates = append(updates, update)
		}
	}
	sum, newupdates := Part1(ordering, updates)
	fmt.Println(sum)
	sum = 0
	for _, newupdate := range newupdates {
		arr := Part2(ordering, newupdate)
		sum += arr[len(arr)/2]
	}
	fmt.Println(sum)
}

func Part2(ordering map[int][]int, update []int) []int {
	changed := true
	current := make([]int, len(update))
	copy(current, update)

	for changed {
		changed = false
		for i := 0; i < len(current)-1; i++ {
			if !Part1Helper(current[i:], ordering[current[i]]) {
				current[i], current[i+1] = current[i+1], current[i]
				changed = true
			}
		}
	}

	return current
}

func Part1(ordering map[int][]int, updates [][]int) (int, [][]int) {
	sum := 0
	newupdates := [][]int{}
	var isvalid bool
	for _, update := range updates {
		for i, page := range update {
			if i != len(update)-1 {

				arr := ordering[page]
				if !Part1Helper(update[i:], arr) {
					newupdates = append(newupdates, update)
					isvalid = false
					break
				}
				isvalid = true
			}
		}
		if isvalid {
			sum += update[len(update)/2]
		}
	}
	return sum, newupdates
}

func Part1Helper(update, arr []int) bool {
	for _, v := range update {
		for _, n := range arr {
			if v == n {
				return false
			}
		}
	}
	return true
}

func strToINt(arr []string) []int {
	slice := []int{}

	for _, v := range arr {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		slice = append(slice, n)
	}
	return slice
}
