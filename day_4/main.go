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
	fmt.Println(Part1(filearr))
}

func Part1(arr []string) int {
	count := 0
	for i, str := range arr {
		for j, c := range str {
			if c == 'X' {
				// forward horizontal
				if j <= len(str)-4 && str[j+1] == 'M' && str[j+2] == 'A' && str[j+3] == 'S' {
					count++
				}
				// forward vertical
				if i <= len(arr)-4 && arr[i+1][j] == 'M' && arr[i+2][j] == 'A' && arr[i+3][j] == 'S' {
					count++
				}
				// backward horizontal
				if j >= 3 && str[j-1] == 'M' && str[j-2] == 'A' && str[j-3] == 'S' {
					count++
				}
				// backward vertical
				if i >= 3 && arr[i-1][j] == 'M' && arr[i-2][j] == 'A' && arr[i-3][j] == 'S' {
					count++
				}
				// forward diagonal
				if j <= len(str)-4 && i <= len(arr)-4 && arr[i+1][j+1] == 'M' && arr[i+2][j+2] == 'A' && arr[i+3][j+3] == 'S' {
					count++
				}
				// backward diagonal
				if i >= 3 && j >= 3 && arr[i-1][j-1] == 'M' && arr[i-2][j-2] == 'A' && arr[i-3][j-3] == 'S' {
					count++
				}
				if i >= 3 && j <= len(str)-4 && arr[i-1][j+1] == 'M' && arr[i-2][j+2] == 'A' && arr[i-3][j+3] == 'S' {
					count++
				}
				if j >= 3 && i <= len(arr)-4 && arr[i+1][j-1] == 'M' && arr[i+2][j-2] == 'A' && arr[i+3][j-3] == 'S' {
					count++
				}

			}
		}
	}
	return count
}