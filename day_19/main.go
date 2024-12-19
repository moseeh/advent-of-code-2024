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
	filearr := strings.Split(filestr, "\n\n")

	towelPatterns := strings.Split(filearr[0], ", ")

	towelDesigns := strings.Split(filearr[1], "\n")
	fmt.Println(len(towelPatterns), len(towelDesigns[10]))
	fmt.Println(Part1(towelPatterns, towelDesigns))
	fmt.Println(Part2(towelPatterns, towelDesigns))

}

func Part1(patterns, designs []string) int {
	count := 0
	for _, str := range designs {
		if canBeDecomposedDP(str, patterns) {
			count++
		}
	}
	return count
}

func canBeDecomposedDP(str string, patterns []string) bool {
	n := len(str)
	dp := make([]bool, n+1)
	dp[n] = true

	for i := n - 1; i >= 0; i-- {
		for _, pattern := range patterns {
			if i+len(pattern) <= n && str[i:i+len(pattern)] == pattern && dp[i+len(pattern)] {
				dp[i] = true
				break
			}
		}
	}

	return dp[0]
}

func Part2(patterns, designs []string) int {
	total := 0
	for _, str := range designs {
		total += countWaysDP(str, patterns)
	}
	return total
}

func countWaysDP(str string, patterns []string) int {
	n := len(str)
	dp := make([]int, n+1)
	dp[n] = 1

	for i := n - 1; i >= 0; i-- {
		for _, pattern := range patterns {
			if i+len(pattern) <= n && str[i:i+len(pattern)] == pattern {
				dp[i] += dp[i+len(pattern)]
			}
		}
	}

	return dp[0]
}
