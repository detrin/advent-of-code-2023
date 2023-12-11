package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	r, c int
}

func NewPos(r, c int) Pos {
	return Pos{r, c}
}

type Graph struct {
	edges map[Pos]map[Pos]struct{}
}

func NewGraph() *Graph {
	return &Graph{edges: make(map[Pos]map[Pos]struct{})}
}

func (g *Graph) AddEdge(u, v Pos) {
	if _, exists := g.edges[u]; !exists {
		g.edges[u] = make(map[Pos]struct{})
	}
	g.edges[u][v] = struct{}{}
}

func (g *Graph) DFS(current Pos, start Pos, visited *map[Pos]bool, mainLoop *map[Pos]bool, depth int) int {
	(*visited)[current] = true
	defer func() { (*visited)[current] = false }()
	maxDepth := depth
	for nbr := range g.edges[current] {
		if nbr == start && maxDepth > 4{
			*mainLoop = make(map[Pos]bool)
			for pos := range *visited {
				(*mainLoop)[pos] = true
			}
			return depth+1
		}
		if !(*visited)[nbr] {
			maxDepth = max(maxDepth, g.DFS(nbr, start, visited, mainLoop, depth+1))
		}
	}
	return maxDepth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isInsideIncrement(pos Pos, loop map[Pos]bool, g *Graph, m int) float64 {
	crosses := 0.0
	if _, inLoop := loop[pos]; !inLoop {
		return 0.0
	}
	_, isBelow := g.edges[pos][NewPos(pos.r-1, pos.c)]
	_, isAbove := g.edges[pos][NewPos(pos.r+1, pos.c)]
	_, isPrev := g.edges[pos][NewPos(pos.r, pos.c-1)]
	_, isNext := g.edges[pos][NewPos(pos.r, pos.c+1)]
	if isBelow && isAbove { // |
		crosses += 1
	} else if isBelow && isNext { // L
		crosses += 0.5
	} else if isBelow && isPrev { // J
		crosses -= 0.5
	} else if isAbove && isNext { // F
		crosses -= 0.5
	} else if isAbove && isPrev { // 7
		crosses += 0.5
	}
	return crosses
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var buffer []string
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)  
		buffer = append(buffer, line)
	}
	m := len(buffer[0])
	n := len(buffer)

	g := NewGraph()
	var animal Pos
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			char := rune(buffer[i][j])
			pos := NewPos(i, j)
			switch char {
			case 'S':
				animal = pos
			case '|':
				g.AddEdge(pos, NewPos(i-1, j))
				g.AddEdge(pos, NewPos(i+1, j))
			case '-':
				g.AddEdge(pos, NewPos(i, j-1))
				g.AddEdge(pos, NewPos(i, j+1))
			case 'L':
				g.AddEdge(pos, NewPos(i-1, j))
				g.AddEdge(pos, NewPos(i, j+1))
			case 'J':
				g.AddEdge(pos, NewPos(i-1, j))
				g.AddEdge(pos, NewPos(i, j-1))
			case '7':
				g.AddEdge(pos, NewPos(i, j-1))
				g.AddEdge(pos, NewPos(i+1, j))
			case 'F':
				g.AddEdge(pos, NewPos(i, j+1))
				g.AddEdge(pos, NewPos(i+1, j))
			}
		}
	}

	deltas := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, delta := range deltas {
		nbr := NewPos(animal.r+delta.r, animal.c+delta.c)
		if _, exists := g.edges[nbr][animal]; exists {
			g.AddEdge(animal, nbr)
		}
	}

	visited := make(map[Pos]bool)
	longestLoop := make(map[Pos]bool)  
	//loopLength := g.DFS(animal, animal, &visited, &longestLoop, 0)
	g.DFS(animal, animal, &visited, &longestLoop, 0)
	
	// task1
	// fmt.Println(loopLength/2) 

	total := 0  
	for i := 0; i < n; i++ {  
		cross_value := 1000.0
		for j := 0; j < m; j++ {  
			pos := NewPos(i, j)  
			increment := isInsideIncrement(pos, longestLoop, g, m)
			cross_value += increment
			if _, inLoop := longestLoop[pos]; !inLoop && int(cross_value) % 2 == 1 {
				total++  
			} 
		}  
	}  
  
	fmt.Println(total)  
}  

