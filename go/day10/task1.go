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

type Board struct {
	grid map[Point][]Point
}

func (b *Board) addPoint(p Point, point Point) {
	if _, ok := b.grid[p]; !ok {
		b.grid[p] = make([]Point, 0)
	}
	b.grid[p] = append(b.grid[p], point)
}

func dfs(b *Board, visited map[Point]bool, curr Point, depth int, start_pt Point) int {  
	visited[curr] = true  
	maxDepth := depth  
	for _, pt := range b.grid[curr] {  
		if pt == start_pt {
			return depth+1
		}
		if !visited[pt] {  
			maxDepth = max(maxDepth, dfs(b, visited, pt, depth+1, start_pt))  
		}  
	}  
	visited[curr] = false
	return maxDepth  
}  


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var buffer []string
	for scanner.Scan() {
		buffer = append(buffer, scanner.Text())
	}
	max_x := len(buffer[0])
	max_y := len(buffer)

	grid := Board{make(map[Point][]Point)}
	curr_pt := Point{0, 0}
	start_pt := Point{0, 0}
	ch := ' '
	for j := 0; j < max_y; j++ {
		curr_pt.x = j
		for i := 0; i < max_x; i++ {
			curr_pt.y = i
			ch = rune(buffer[j][i])
			if ch == 'S' {
				start_pt = curr_pt
			} else if ch != '.' {
				rel_pts := getRelativePoints(ch)
				pt1 := curr_pt.add(rel_pts[0])
				pt2 := curr_pt.add(rel_pts[1])
				if pt1.x >= 0 && pt1.x < max_x && pt1.y >= 0 && pt1.y < max_y {
					grid.addPoint(pt1, curr_pt)
				}
				if pt2.x >= 0 && pt2.x < max_x && pt2.y >= 0 && pt2.y < max_y {
					grid.addPoint(pt2, curr_pt)
				}
			}
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading from stdin:", scanner.Err())
	}
 
    visited := make(map[Point]bool)  
    maxDepth := 0  
    for _, pt := range grid.grid[start_pt] {  
        maxDepth = max(maxDepth, dfs(&grid, visited, pt, 1, start_pt))+1
    } 

	fmt.Println(maxDepth/2)
}
