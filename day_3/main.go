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
	filearr := strings.Split(filestr, "mul")
	fmt.Println(Part1(filearr))
}

func Part1(arr []string) int {
	sum := 0
	for _, v := range arr {
		if v[0] != '(' {
			continue
		}
		if !strings.Contains(v, ")") {
			continue
		}
		b := v[1:strings.Index(v, ")")]
		barr := strings.Split(b, ",")
		if len(barr) != 2 {
			continue
		}
		num1, err := strconv.Atoi(barr[0])
		if err != nil {
			continue
		}
		num2, err := strconv.Atoi(barr[1])
		if err != nil {
			continue
		}
		if num1 > 999 || num2 > 999 {
			continue
		}
		sum += num1 * num2
	}
	return sum
}
