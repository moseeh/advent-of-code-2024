package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Arr []string

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filestr := string(filebytes)
	filearr := strings.Fields(filestr)
	blinks := 1
	Arr = filearr
	part1(Arr, blinks)
	fmt.Println(len(Arr))
}

func part1(arr []string, blinks int) {
	newarr := []string{}
	if blinks > 25 {
		return
	}
	fmt.Println(blinks)
	for _, v := range arr {
		n, _ := strconv.Atoi(v)
		if n == 0 {
			newarr = append(newarr, "1")
		} else {
			if len(v)%2 == 0 {
				l := len(v) / 2
				n1, _ := strconv.Atoi(v[0:l])
				n2, _ := strconv.Atoi(v[l:])
				newarr = append(newarr, strconv.Itoa(n1))
				newarr = append(newarr, strconv.Itoa(n2))
			} else {
				newarr = append(newarr, strconv.Itoa(n*2024))
			}
		}
	}
	blinks++
	Arr = newarr
	part1(Arr, blinks)
}
