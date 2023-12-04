package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	red   int
	green int
	blue  int
}

func parseCubes(s string) Cube {
	parts := strings.Split(s, ", ")
	c := Cube{}
	for _, part := range parts {
		vals := strings.Split(part, " ")
		num, _ := strconv.Atoi(vals[0])
		switch vals[1] {
		case "red":
			c.red += num
		case "green":
			c.green += num
		case "blue":
			c.blue += num
		}
	}
	return c
}

func main() {
	game_mult_sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		sets := strings.Split(parts[1], "; ")
		red_needed := 0
		green_needed := 0
		blue_needed := 0
		for _, set := range sets {
			cube := parseCubes(set)
			if red_needed < cube.red {
				red_needed = cube.red
			}
			if green_needed < cube.green {
				green_needed = cube.green
			}
			if blue_needed < cube.blue {
				blue_needed = cube.blue
			}
		}
		game_mult_sum += red_needed * green_needed * blue_needed
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", scanner.Err())
	}
	fmt.Println(game_mult_sum)
}
