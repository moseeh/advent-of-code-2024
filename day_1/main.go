package main

import (
	"fmt"
	"os"
	"sort"
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
	arrleft, arrright := []int{}, []int{}
	for _, line := range filearr {
		linearr := strings.Fields(line)
		num1, _ := strconv.Atoi(linearr[0])
		num2, _ := strconv.Atoi(linearr[1])
		arrleft = append(arrleft, num1)
		arrright = append(arrright, num2)
	}
	sort.Ints(arrleft)
	sort.Ints(arrright)
	fmt.Println(Part1(arrleft, arrright))
}

func Part1(arr1, arr2 []int) int {
	sum := 0
	for i := 0; i <= len(arr1)-1; i++ {
		add := arr1[i] - arr2[i]
		sum += abs(add)
	}
	return sum
}

func abs(num int) int {
	if num < 0 {
		return -(num)
	}
	return num
}
