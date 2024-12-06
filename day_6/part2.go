package main

import "fmt"

func Part2(area []string) {
	for i := 0; i <= len(area)-1; i++ {
		for j := 0; j <= len(area[i])-1; j++ {
			newarea := []string{}
			for row, line := range area {
				str := ""
				for column, char := range line {
					if char != '#' && i == row && j == column && char != '^' {
						str += "#"
					} else {
						str += string(char)
					}
				}
				newarea = append(newarea, str)
			}
			Solved = false
			for k := range Up {
				Up[k] = false
			}
			for k := range Down {
				Down[k] = false
			}
			for k := range Right {
				Right[k] = false
			}
			for k := range Left {
				Left[k] = false
			}
			for k := range ObstacleMet {
				ObstacleMet[k] = false
			}

			Part1(newarea)
			if !Solved {
				fmt.Println()

				Positions++
			}
		}
	}
}
