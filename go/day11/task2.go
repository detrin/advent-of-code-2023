package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Point struct {
	r, c int
}

func (p Point) add(other Point) Point {
	return Point{p.r + other.r, p.c + other.c}
}

func (p Point) manhattanDistance(other Point) int {
	return abs(p.r-other.r) + abs(p.c-other.c)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var buffer []string
	for scanner.Scan() {
		buffer = append(buffer, scanner.Text())
	}
	max_c := len(buffer[0])
	max_r := len(buffer)

	space_warping_r := make([]int, max_r)
	space_warping_c := make([]int, max_c)
	expanding_factor := 1000000
	for r := 0; r < max_r; r++ {
		empty := true
		for c := 0; c < max_c; c++ {
			if buffer[r][c] == '#' {
				empty = false
				break
			}
		}
		if r > 0 {
			space_warping_r[r] = space_warping_r[r-1]
		}
		if empty {
			space_warping_r[r] += expanding_factor-1
		}
	}

	for c := 0; c < max_c; c++ {
		empty := true
		for r := 0; r < max_r; r++ {
			if buffer[r][c] == '#' {
				empty = false
				break
			}
		}
		if c > 0 {
			space_warping_c[c] = space_warping_c[c-1]
		}
		if empty {
			space_warping_c[c] += expanding_factor-1
		}
	}

	galaxies := make(map[int]Point)
	galaxy_id := 0
	for r := 0; r < max_r; r++ {
		for c := 0; c < max_c; c++ {
			if buffer[r][c] == '#' {
				galaxies[galaxy_id] = Point{r+space_warping_r[r], c+space_warping_c[c]}
				galaxy_id += 1
			}
		}
	}

	total_dist := 0
	for i := 0; i < galaxy_id-1; i++ {
		for j := i+1; j < galaxy_id; j++ {
			total_dist += galaxies[i].manhattanDistance(galaxies[j])
		}
	}

	fmt.Println(total_dist)
}