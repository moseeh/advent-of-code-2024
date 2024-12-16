package main

import "fmt"

func part2(final map[[2]int]int) {
	fmt.Println(count)
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			if final[[2]int{j, i}] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(final[[2]int{j, i}])
			}
		}
		fmt.Println()
	}
}
