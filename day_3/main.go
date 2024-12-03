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
	fmt.Println(Part2(filestr))
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

func Part2(str string) int {
	sum := 0
	do := true
	start := 0
	for start < len(str) {
		if strings.HasPrefix(str[start:], "don't()") {
			do = false
			start += len("don't()")
			continue
		}
		if strings.HasPrefix(str[start:], "do()") {
			do = true
			start += len("do()")
			continue
		}
		if do {
			if strings.HasPrefix(str[start:], "mul(") {
				s := str[start+len("mul("):]
				s = s[:strings.Index(s, ")")]
				sarr := strings.Split(s, ",")
				if len(sarr) != 2 {
					start++
					continue
				}
				num1, err := strconv.Atoi(sarr[0])
				if err != nil {
					start++
					continue
				}
				num2, err := strconv.Atoi(sarr[1])
				if err != nil {
					start++
					continue
				}
				if num1 > 999 || num2 > 999 {
					start++
					continue
				}
				sum += num1 * num2
			}
		}
		start++
	}
	return sum
}
