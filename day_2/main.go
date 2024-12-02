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
	var count int
	for _, line := range filearr {
		arr := []int{}
		linearr := strings.Fields(line)
		for _, v := range linearr {
			num, _ := strconv.Atoi(v)
			arr = append(arr, num)
		}
		if Part1(arr) {
			fmt.Println(arr)
			count++
		}
	}
	fmt.Println(count)
}

func Part1(arr []int) bool {
	if arr[0] >= arr[1] {
		for i := 1; i < len(arr); i++ {
			if arr[i-1]-arr[i] < 1 || arr[i-1]-arr[i] > 3 {
				return false
			}
		}
	}
	if arr[0] <= arr[1] {
		for i := 1; i < len(arr); i++ {
			if arr[i]-arr[i-1] < 1 || arr[i]-arr[i-1] > 3 {
				return false
			}
		}
	}
	return true
}
