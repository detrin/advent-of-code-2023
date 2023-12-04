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
	red_limit := 12
	green_limit := 13
	blue_limit := 14
	game_id_sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		game_id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
		sets := strings.Split(parts[1], "; ")
		all_games_possible := true
		for _, set := range sets {
			cube := parseCubes(set)
			all_games_possible = all_games_possible && cube.red <= red_limit && cube.green <= green_limit && cube.blue <= blue_limit
		}
		if all_games_possible {
			game_id_sum += game_id
		}
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", scanner.Err())
	}
	fmt.Println(game_id_sum)
}
