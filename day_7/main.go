package main

import (
	"fmt"
	"math"
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
	sum := 0
	for _, line := range filearr {
		sum += Part1(line)
	}
	fmt.Println(sum)
}

func Part1(line string) int {
	linearr := strings.Split(line, ":")

	value, _ := strconv.Atoi(linearr[0])
	arr := strings.Fields(linearr[1])
	numarr := []int{}
	for _, str := range arr {
		num, _ := strconv.Atoi(str)
		numarr = append(numarr, num)
	}
	if Part1helper(value, numarr) {
		return value
	}
	return 0
}

func Part1helper(num int, numarr []int) bool {
	tries := math.Pow(2.0, float64(len(numarr)-1))
	for i := 0; i <= int(tries)-1; i++ {
		str := strconv.FormatInt(int64(i), 2)

		if len(str) != len(numarr)-1 {
			str = strings.Repeat("0", len(numarr)-1-len(str)) + str
		}
		sum := 0
		for i, v := range str {
			if sum == 0 {
				sum = numarr[i]
			}
			if v == '0' {
				sum *= numarr[i+1]
			} else {
				sum += numarr[i+1]
			}
		}
		if sum == num {
			return true
		}
	}

	return false
}
