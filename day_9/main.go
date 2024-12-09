package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tried = make(map[string]bool)

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filestr := string(filebytes)
	fmt.Println(Part1(filestr))
}

func Part1(filestr string) int {
	arr := []string{}
	mymap := make(map[string]int)
	count := 0
	for i, c := range filestr {
		n, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			arr = append(arr, strconv.Itoa(count))
			mymap[strconv.Itoa(count)] = n
			count++
		} else {
			arr = append(arr, strings.Repeat(".", n))
		}
	}
	// newarr := part1helper(arr, mymap)
	newarr2 := Part2(arr, mymap)
	sum, sum1 := 0, 0
	// for i, v := range newarr {
	// 	n, _ := strconv.Atoi(v)
	// 	sum += i * n
	// }
	for i, v := range newarr2 {
		n, err := strconv.Atoi(v)
		if err == nil {
			sum1 += i * n
		}
	}

	fmt.Println(sum1)

	return sum
}

func Part2(arr []string, mymap map[string]int) []string {
	newarr := []string{}

	for i, v := range arr {
		if i%2 == 0 {
			if !tried[v] {
				for i := 1; i <= mymap[v]; i++ {
					newarr = append(newarr, v)
				}
				tried[v] = true
			} else {
				for mymap[v] > 0 {
					newarr = append(newarr, ".")
					mymap[v]--
				}
			}
		} else {
			if len(v) == 0 {
				continue
			}
			newarr = append(newarr, part2helper(v, arr, mymap)...)
		}
	}
	return newarr
}

func part2helper(str string, arr []string, mymap map[string]int) []string {
	newarr := []string{}
	l := len(str)
	for i := len(arr) - 1; i >= 0; i-- {
		if i%2 == 0 {
			if l < mymap[arr[i]] {
				continue
			}
			if tried[arr[i]] {
				continue
			}
			l = l - mymap[arr[i]]
			freq := mymap[arr[i]]
			for freq > 0 {
				newarr = append(newarr, arr[i])
				freq--
			}
			tried[arr[i]] = true
		}
	}
	for l > 0 {
		newarr = append(newarr, ".")
		l--
	}
	return newarr
}

func part1helper(arr []string, mymap map[string]int) []string {
	newarr := []string{}

	for i, v := range arr {
		if i%2 == 0 && mymap[v] > 0 {
			for i := 1; i <= mymap[v]; i++ {
				newarr = append(newarr, v)
			}
		} else {
			l := len(v)
			if l == 0 {
				continue
			}
			if mymap[arr[i+1]] == 0 {
				break
			}
		outerloop:
			for j := len(arr) - 1; j >= 0; j-- {
				if j%2 == 0 {
					if mymap[arr[j]] == 0 {
						continue
					}
					for mymap[arr[j]] > 0 {
						if l == 0 {
							break outerloop
						}
						newarr = append(newarr, arr[j])
						mymap[arr[j]]--
						l--
					}
				}
			}
		}
	}
	return newarr
}
