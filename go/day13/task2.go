package main

import (
	"bufio"
	"fmt"
	"os"
)

func diffHorizontal(grid []string, i int) int {
	dist := min(i, len(grid)-i)
	sum := 0
	for k := 0; k < len(grid[0]); k++ {
		for j := 0; j < dist; j++ {
			if grid[i+j][k] != grid[i-j-1][k] {
				sum++
			}
		}
	}
	return sum
}

func totalHorizontalDiffs(grid []string, flips int) int {
	sum := 0
	for i := 1; i < len(grid); i++ {
		if diffHorizontal(grid, i) == flips {
			sum += i
		}
	}
	return sum
}

func checkMirror(stdin *bufio.Scanner, flips int) int {
	h, v := 0, 0
	var grid []string
	for stdin.Scan() {
		line := stdin.Text()
		if line == "" {
			h += totalHorizontalDiffs(grid, flips)
			grid = transpose(grid)
			v += totalHorizontalDiffs(grid, flips)
			grid = []string{}
		} else {
			grid = append(grid, line)
		}
	}
	h += totalHorizontalDiffs(grid, flips)
	grid = transpose(grid)
	v += totalHorizontalDiffs(grid, flips)
	return h*100 + v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func transpose(grid []string) []string {
	transposed := make([]string, len(grid[0]))
	for i := range transposed {
		for _, row := range grid {
			transposed[i] += string(row[i])
		}
	}
	return transposed
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	part2 := checkMirror(scanner, 1)
	fmt.Println(part2)
}
