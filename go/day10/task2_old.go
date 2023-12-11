package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func (p Point) add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func getRelativePoints(r rune) [2]Point {
	switch r {
	case '|':
		return [2]Point{{-1, 0}, {1, 0}}
	case '-':
		return [2]Point{{0, -1}, {0, 1}}
	case 'L':
		return [2]Point{{-1, 0}, {0, 1}}
	case 'J':
		return [2]Point{{-1, 0}, {0, -1}}
	case '7':
		return [2]Point{{1, 0}, {0, -1}}
	case 'F':
		return [2]Point{{1, 0}, {0, 1}}
	default:
		return [2]Point{{0, 0}, {0, 0}}
	}
}

func getRune(pts [2]Point) rune {
	switch pts {
	case [2]Point{{-1, 0}, {1, 0}}:
		return '|'
	case [2]Point{{1, 0}, {-1, 0}}:
		return '|'
	case [2]Point{{0, -1}, {0, 1}}:
		return '-'
	case [2]Point{{0, 1}, {0, -1}}:
		return '-'
	case [2]Point{{-1, 0}, {0, 1}}:
		return 'L'
	case [2]Point{{0, 1}, {-1, 0}}:
		return 'L'
	case [2]Point{{-1, 0}, {0, -1}}:
		return 'J'
	case [2]Point{{0, -1}, {-1, 0}}:
		return 'J'
	case [2]Point{{1, 0}, {0, -1}}:
		return '7'
	case [2]Point{{0, -1}, {1, 0}}:
		return '7'
	case [2]Point{{1, 0}, {0, 1}}:
		return 'F'
	case [2]Point{{0, 1}, {1, 0}}:
		return 'F'
	default:
		return ' '
	}
}

func getRelativeBoundaries(r rune) float64 {
	switch r {
	case '|':
		return 1
	case 'L', '7':
		return 0.5
	case 'J', 'F':
		return -0.5
	default:
		return 0
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var buffer []string
	for scanner.Scan() {
		buffer = append(buffer, scanner.Text())
	}
	max_x := len(buffer[0])
	max_y := len(buffer)

	start_pt := Point{0, 0}
	for j := 0; j < max_y; j++ {
		for i := 0; i < max_x; i++ {
			if buffer[j][i] == 'S' {
				start_pt = Point{i, j}
			} 
		}
	}

	directions := [4]Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	selected_directions := make([]Point, 0)
	for _, direction := range directions {
		pt := start_pt.add(direction)
		if pt.x >= 0 && pt.x < max_x && pt.y >= 0 && pt.y < max_y {
			possible_new_dirs := getRelativePoints(rune(buffer[pt.y][pt.x]))
			new_pt := pt.add(possible_new_dirs[0])
			if new_pt == start_pt {
				selected_directions = append(selected_directions, direction)
			}
			new_pt = pt.add(possible_new_dirs[1])
			if new_pt == start_pt {
				selected_directions = append(selected_directions, direction)
			}
		}
	}
	start_rune := getRune([2]Point{selected_directions[0], selected_directions[1]})

	grid := make([][]float64, max_y)
	for i := range grid {
		grid[i] = make([]float64, max_x)
	}
	ch := ' '
	for j := 0; j < max_y; j++ {
		for i := 0; i < max_x; i++ {
			if buffer[j][i] == 'S' {
				ch = start_rune
				ch = 'J'
			} else {
				ch = rune(buffer[j][i])
			}
			grid[j][i] = getRelativeBoundaries(ch)
			
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading from stdin:", scanner.Err())
	}

	insideCount := 0
	inside := make([][]bool, 0)
	for j := 0; j < max_y; j++ {
		inside = append(inside, make([]bool, max_x))
		count := 0.0
		for i := 0; i < max_x; i++ {
			count += grid[j][i]
			is_inside := int(count)%2 != 0 && buffer[j][i] == '.'
			if is_inside {
				insideCount++
			}
			inside[j][i] = is_inside
		}
	}

	for i := 0; i < max_y; i++ {
		for j := 0; j < max_x; j++ {
			if inside[i][j] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println(insideCount)
}
