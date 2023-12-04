package main

import (
	"bufio"
	"fmt"
	"os"
)

func is_number(r rune) bool {
	return '0' <= r && r <= '9'
}

func is_symbol(r rune) bool {
	return !is_number(r) && r != '.'
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", scanner.Err())
	}

	width := len(grid[0])
	height := len(grid)

	within_symbol := make([][]bool, height)
	for i := range within_symbol {
		within_symbol[i] = make([]bool, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if is_symbol(grid[y][x]) {
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if nx, ny := x+dx, y+dy; 0 <= nx && nx < width && 0 <= ny && ny < height {
							within_symbol[ny][nx] = true
						}
					}
				}
			}
		}
	}

	numbers_within_symbol := 0
	for y := 0; y < height; y++ {
		number := 0
		is_within_symbol := false
		for x := 0; x < width; x++ {
			if is_number(grid[y][x]) {
				number = number*10 + int(grid[y][x]-'0')
				is_within_symbol = is_within_symbol || within_symbol[y][x]
			} else {
				if is_within_symbol {
					numbers_within_symbol += number
				}
				number = 0
				is_within_symbol = false
			}
		}
		if is_within_symbol {
			numbers_within_symbol += number
		}
	}

	fmt.Println(numbers_within_symbol)
}
