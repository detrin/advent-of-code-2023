package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cards_pts_sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		num_parts := strings.Split(parts[1], " | ")
		winning_nums := make([]int, 0)
		for _, num := range strings.Split(num_parts[0], " ") {
			if num == "" {
				continue
			}
			n, _ := strconv.Atoi(num)
			winning_nums = append(winning_nums, n)
		}
		pts := 0
		for _, num := range strings.Split(num_parts[1], " ") {
			actual_num, _ := strconv.Atoi(num)
			for _, winning_num := range winning_nums {
				if winning_num == actual_num {
					if pts == 0 {
						pts = 1
					} else {
						pts *= 2
					}
				}
			}
		}

		cards_pts_sum += pts

	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", scanner.Err())
	}
	fmt.Println(cards_pts_sum)
}
