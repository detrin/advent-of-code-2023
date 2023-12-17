package main

import (
	"bufio"
	"fmt"
	"os"
)

func rotate(G [][]rune) [][]rune {
	R := len(G)
	C := len(G[0])
	NG := make([][]rune, C)
	for i := range NG {
		NG[i] = make([]rune, R)
	}
	for r := range G {
		for c := range G[r] {
			NG[c][R-1-r] = G[r][c]
		}
	}
	return NG
}

func roll(G [][]rune) [][]rune {
	R := len(G)
	C := len(G[0])
	for c := 0; c < C; c++ {
		for _ = range G {
			for r := 0; r < R; r++ {
				if G[r][c] == 'O' && r > 0 && G[r-1][c] == '.' {
					G[r][c] = '.'
					G[r-1][c] = 'O'
				}
			}
		}
	}
	return G
}

func score(G [][]rune) int {
	ans := 0
	R := len(G)
	C := len(G[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if G[r][c] == 'O' {
				ans += len(G) - r
			}
		}
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	G := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		G = append(G, []rune(line))
	}

	byGrid := make(map[string]int)
	target := 1000000000
	t := 0
	for t < target {
		t++
		for j := 0; j < 4; j++ {
			G = roll(G)
			G = rotate(G)
		}
		Gh := fmt.Sprint(G)
		if val, ok := byGrid[Gh]; ok {
			cycle_length := t - val
			amt := (target - t) / cycle_length
			t += amt * cycle_length
		}
		byGrid[Gh] = t
	}
	fmt.Println(score(G))
}
