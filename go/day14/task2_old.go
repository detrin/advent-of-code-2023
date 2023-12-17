package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid struct {
	Rows [][]int
}

func NewGrid() *Grid {
	return &Grid{
		Rows: make([][]int, 0),
	}
}

func (g *Grid) AddRow(row []int) {
	g.Rows = append(g.Rows, row)
}

func (g *Grid) MoveToLeft() {
	for i := range g.Rows {
		g.sortSubline(i)
	}
}

func (g *Grid) sortSubcells(start, end, index int) {
	movable := 0
	for i := start; i < end; i++ {
		if g.Rows[index][i] == 2 {
			movable++
		}
	}
	for i := start; i < end; i++ {
		if movable > 0 {
			g.Rows[index][i] = 2
			movable--
		} else {
			g.Rows[index][i] = 0
		}
	}
}

func (g *Grid) sortSubline(index int) {
	row := &g.Rows[index]
	start := 0
	for i, cell := range *row {
		if cell == 1 {
			g.sortSubcells(start, i, index)
			start = i + 1
		}
	}
	g.sortSubcells(start, len(*row), index)
}

func (g *Grid) RotateClockwise() {
	n := len(g.Rows)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			g.Rows[i][j], g.Rows[j][i] = g.Rows[j][i], g.Rows[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			g.Rows[i][j], g.Rows[i][k] = g.Rows[i][k], g.Rows[i][j]
		}
	}
}

func (g *Grid) RotateAntiClockwise() {
	n := len(g.Rows)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			g.Rows[i][j], g.Rows[j][i] = g.Rows[j][i], g.Rows[i][j]
		}
	}

	for j := 0; j < n; j++ {
		for i, k := 0, n-1; i < k; i, k = i+1, k-1 {
			g.Rows[i][j], g.Rows[k][j] = g.Rows[k][j], g.Rows[i][j]
		}
	}
}

func (g *Grid) Rotate180() {
	n := len(g.Rows)
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			g.Rows[i][j], g.Rows[i][k] = g.Rows[i][k], g.Rows[i][j]
		}
	}

	for j := 0; j < n; j++ {
		for i, k := 0, n-1; i < k; i, k = i+1, k-1 {
			g.Rows[i][j], g.Rows[k][j] = g.Rows[k][j], g.Rows[i][j]
		}
	}
}

func (g *Grid) Println() {
	for _, row := range g.Rows {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else if cell == 1 {
				fmt.Print("#")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func (g *Grid) CalculateLoad() int {
	totalLoad := 0
	n := len(g.Rows)

	for i := range g.Rows {
		for _, cell := range g.Rows[i] {
			if cell == 2 {
				totalLoad += n - i
			}
		}
	}

	return totalLoad
}

func (g *Grid) ToString() string {
	var str string
	for _, row := range g.Rows {
		for _, cell := range row {
			if cell == 0 || cell == 2 {
				str += fmt.Sprintf("%d", cell)
			}
		}
	}
	return str
}

func (g *Grid) Hash() uint64 {
	offset64 := uint64(14695981039346656037)
	prime64 := uint64(1099511628211)

	str := g.ToString()
	hash := offset64
	for i := 0; i < len(str); i++ {
		hash ^= uint64(str[i])
		hash *= prime64
	}
	return hash
}

func (g *Grid) Cycle() {
	g.MoveToLeft()
	g.RotateClockwise()
	g.MoveToLeft()
	g.RotateClockwise()
	g.MoveToLeft()
	g.RotateClockwise()
	g.MoveToLeft()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := NewGrid()

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, ch := range line {
			if ch == '.' {
				row = append(row, 0)
			} else if ch == '#' {
				row = append(row, 1)
			} else {
				row = append(row, 2)
			}
		}
		grid.AddRow(row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	step_records := make(map[uint64]int)
	grid.RotateAntiClockwise()
	steps := 0
	for steps < 1000000000 {
		grid.Cycle()
		grid.RotateClockwise()

		load := grid.CalculateLoad()
		hash := grid.Hash()
		grid.RotateClockwise()
		fmt.Println("steps", steps, load)
		grid.Println()
		fmt.Println()
		grid.RotateAntiClockwise()

		if load_stored, found := step_records[hash]; found {
			loop_length := steps - load_stored
			steps += loop_length * ((1000000000 - steps) / loop_length)
		}
		step_records[hash] = steps
		steps++
	}
	load := grid.CalculateLoad()
	grid.RotateClockwise()
	grid.Println()
	fmt.Println(load)
}
