package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func processGear(gearPointer []int, number int, numbersPerGear, gearRatio []int) {
	for _, gear := range gearPointer {
		numbersPerGear[gear-1]++
		gearRatio[gear-1] *= number
	}
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var grid [][]string
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", scanner.Err())
	}

	width := len(grid[0])
	height := len(grid)

	gearsTotal := 0
	gearsPointers := make([][]int, height)
	for i := range gearsPointers {
		gearsPointers[i] = make([]int, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == "*" {
				gearsTotal++
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if x+dx >= 0 && x+dx < width && y+dy >= 0 && y+dy < height {
							gearsPointers[y+dy][x+dx] = gearsTotal
						}
					}
				}
			}
		}
	}

	numbersPerGear := make([]int, gearsTotal)
	gearRatio := make([]int, gearsTotal)
	for i := range gearRatio {
		gearRatio[i] = 1
	}

	for y := 0; y < height; y++ {
		number := 0
		gearPointer := make([]int, 0)
		for x := 0; x < width; x++ {
			if isNumber(rune(grid[y][x][0])) {
				number = number*10 + int(grid[y][x][0]-'0')
				if gearsPointers[y][x] != 0 && !contains(gearPointer, gearsPointers[y][x]) {
					gearPointer = append(gearPointer, gearsPointers[y][x])
				}
			} else {
				processGear(gearPointer, number, numbersPerGear, gearRatio)
				gearPointer = make([]int, 0)
				number = 0
			}
		}
		processGear(gearPointer, number, numbersPerGear, gearRatio)
	}

	gearRatioSum := 0
	for i := range gearRatio {
		if numbersPerGear[i] == 2 {
			gearRatioSum += gearRatio[i]
		}
	}
	fmt.Println(gearRatioSum)
}
