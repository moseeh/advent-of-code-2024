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
	fmt.Println(Part1(ordering, updates))
}

func Part1(ordering map[int][]int, updates [][]int) int {
	sum := 0

	var isvalid bool
	for _, update := range updates {
		for i, page := range update {
			if i != len(update)-1 {

				arr := ordering[page]
				if !Part1Helper(update[i:], arr) {
					isvalid = false
					break
				}
				isvalid = true
			}
		}
		if isvalid {
			fmt.Println(update)
			sum += update[len(update)/2]
		}
	}
	return sum
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
