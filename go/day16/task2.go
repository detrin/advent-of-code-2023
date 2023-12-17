package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Node struct {
	value     rune
	energized bool
}

type NodeKey struct {
	x, y int
	dir  Direction
}

type Graph struct {
	nodes   [][]Node
	visited map[NodeKey]bool
}

func NewGraph(lines []string) *Graph {
	height := len(lines)
	width := len(lines[0])
	nodes := make([][]Node, height)
	for i := range nodes {
		nodes[i] = make([]Node, width)
		for j := range nodes[i] {
			nodes[i][j] = Node{rune(lines[i][j]), false}
		}
	}
	return &Graph{nodes, make(map[NodeKey]bool)}
}

func (g *Graph) Energize(x, y int, dir Direction) {
	if x < 0 || y < 0 || x >= len(g.nodes[0]) || y >= len(g.nodes) {
		return
	}
	key := NodeKey{x, y, dir}
	if _, ok := g.visited[key]; ok {
		return
	}
	g.visited[key] = true
	node := &g.nodes[y][x]
	node.energized = true
	switch node.value {
	case '.':
		g.moveStraight(x, y, dir)
	case '/':
		g.moveReflected(x, y, dir, '/')
	case '\\':
		g.moveReflected(x, y, dir, '\\')
	case '|', '-':
		g.splitBeam(x, y, dir, node.value)
	}
}

func (g *Graph) moveStraight(x, y int, dir Direction) {
	switch dir {
	case Up:
		g.Energize(x, y-1, Up)
	case Down:
		g.Energize(x, y+1, Down)
	case Left:
		g.Energize(x-1, y, Left)
	case Right:
		g.Energize(x+1, y, Right)
	}
}

func (g *Graph) moveReflected(x, y int, dir Direction, mirror rune) {
	switch dir {
	case Up:
		if mirror == '/' {
			g.Energize(x+1, y, Right)
		} else {
			g.Energize(x-1, y, Left)
		}
	case Down:
		if mirror == '/' {
			g.Energize(x-1, y, Left)
		} else {
			g.Energize(x+1, y, Right)
		}
	case Left:
		if mirror == '/' {
			g.Energize(x, y+1, Down)
		} else {
			g.Energize(x, y-1, Up)
		}
	case Right:
		if mirror == '/' {
			g.Energize(x, y-1, Up)
		} else {
			g.Energize(x, y+1, Down)
		}
	}
}

func (g *Graph) splitBeam(x, y int, dir Direction, splitter rune) {
	switch dir {
	case Up:
		if splitter == '-' {
			g.Energize(x+1, y, Right)
			g.Energize(x-1, y, Left)
		} else {
			g.Energize(x, y-1, Up)
		}
	case Down:
		if splitter == '-' {
			g.Energize(x+1, y, Right)
			g.Energize(x-1, y, Left)
		} else {
			g.Energize(x, y+1, Down)
		}
	case Left:
		if splitter == '|' {
			g.Energize(x, y+1, Down)
			g.Energize(x, y-1, Up)
		} else {
			g.Energize(x-1, y, Left)
		}
	case Right:
		if splitter == '|' {
			g.Energize(x, y+1, Down)
			g.Energize(x, y-1, Up)
		} else {
			g.Energize(x+1, y, Right)
		}
	}
}

func (g *Graph) CountEnergized() int {
	count := 0
	for _, row := range g.nodes {
		for _, node := range row {
			if node.energized {
				count++
			}
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	size_x := len(lines[0])
	size_y := len(lines)

	// make a grapph, then iterate over bounders and energize
	// save the max energized
	energized_max := 0
	for x := 0; x < size_x; x++ {
		graph := NewGraph(lines)
		graph.Energize(x, 0, Down)
		energized := graph.CountEnergized()
		energized_max = max(energized_max, energized)

		graph = NewGraph(lines)
		graph.Energize(x, size_y-1, Up)
		energized = graph.CountEnergized()
		energized_max = max(energized_max, energized)
	}

	for y := 0; y < size_y; y++ {
		graph := NewGraph(lines)
		graph.Energize(0, y, Right)
		energized := graph.CountEnergized()
		energized_max = max(energized_max, energized)

		graph = NewGraph(lines)
		graph.Energize(size_x-1, y, Left)
		energized = graph.CountEnergized()
		energized_max = max(energized_max, energized)
	}

	fmt.Println(energized_max)
}
