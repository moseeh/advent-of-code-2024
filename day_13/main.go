package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	ButtonA [2]int
	ButtonB [2]int
	Prize   [2]int
}

func main() {
	filebytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filestr := string(filebytes)
	filearr := strings.Split(filestr, "\n\n")
	var machines []Machine

	for _, v := range filearr {
		var machine Machine

		lines := strings.Split(v, "\n")
		for _, part := range lines {
			parts := strings.Split(part, ":")
			if parts[0] == "Button A" {
				axis := strings.Split(parts[1], ", ")
				x, _ := strconv.Atoi(axis[0][2:])
				y, _ := strconv.Atoi(axis[1][1:])
				machine.ButtonA = [2]int{x, y}
			}
			if parts[0] == "Button B" {
				axis := strings.Split(parts[1], ", ")
				x, _ := strconv.Atoi(axis[0][2:])
				y, _ := strconv.Atoi(axis[1][1:])
				machine.ButtonB = [2]int{x, y}
			}
			if parts[0] == "Prize" {
				axis := strings.Split(parts[1], ", ")
				x, _ := strconv.Atoi(axis[0][3:])
				y, _ := strconv.Atoi(axis[1][2:])
				machine.Prize = [2]int{x, y}
			}
		}
		machines = append(machines, machine)
	}
	tokens := 0
	for _, machine := range machines {
		tokens += part1(machine)
	}
	fmt.Println(tokens)
}

func part1(machine Machine) int {
	A, B := SolveEquations(machine.ButtonA[0], machine.ButtonB[0], machine.Prize[0], machine.ButtonA[1], machine.ButtonB[1], machine.Prize[1])
	return A*3 + B*1
}

// SolveEquations solves a system of two linear equations:
// c1a + c2b = c3
// c4a + c5b = c6
func SolveEquations(c1, c2, c3, c4, c5, c6 int) (int, int) {
	// Calculate the determinant of the system
	det := c1*c5 - c2*c4
	if det == 0 {
		return 0, 0
	}
	detA := c3*c5 - c2*c6
	detB := c1*c6 - c3*c4

	// Check if determinants are divisible by det
	if detA%det != 0 || detB%det != 0 {
		return 0, 0
	}

	// Use the elimination method to solve for a and b
	a := (c3*c5 - c2*c6) / det
	b := (c1*c6 - c3*c4) / det

	return a, b
}
