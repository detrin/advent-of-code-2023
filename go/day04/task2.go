package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cards struct {
	count []int
}

func assureCapacity(c *Cards, capacity int) {
	if len(c.count) < capacity {
		for i := len(c.count); i < capacity; i++ {
			c.count = append(c.count, 1)
		}
	}
}

func addCards(c *Cards, pos int, count int) {
	assureCapacity(c, pos+count+1)
	for i := pos + 1; i < pos+count+1; i++ {
		c.count[i] += c.count[pos]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	generated_cards := Cards{}
	card_num := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		// card_num, _ = strconv.Atoi(strings.Split(parts[0], " ")[1])
		num_parts := strings.Split(parts[1], " | ")
		winning_nums := make([]int, 0)
		for _, num := range strings.Split(num_parts[0], " ") {
			if num == "" {
				continue
			}
			n, _ := strconv.Atoi(num)
			winning_nums = append(winning_nums, n)
		}
		actual_nums := make([]int, 0)
		for _, num := range strings.Split(num_parts[1], " ") {
			if num == "" {
				continue
			}
			n, _ := strconv.Atoi(num)
			actual_nums = append(actual_nums, n)
		}

		assureCapacity(&generated_cards, card_num+1)
		matched := 0
		for _, num := range strings.Split(num_parts[1], " ") {
			actual_num, _ := strconv.Atoi(num)
			for _, winning_num := range winning_nums {
				if winning_num == actual_num {
					matched++
				}
			}
		}
		if matched > 0 {
			addCards(&generated_cards, card_num, matched)
		}
		card_num++

	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", scanner.Err())
	}
	generated_cards_total := 0
	for i := 0; i < card_num; i++ {
		generated_cards_total += generated_cards.count[i]
	}
	fmt.Println(generated_cards_total)
}
