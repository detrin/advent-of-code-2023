package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cacheKey struct {
	lava    string
	springs string
}

var cache = make(map[cacheKey]int)

func recurse(lava string, springs []int) int {
	key := cacheKey{lava, fmt.Sprint(springs)}
	if _, found := cache[key]; found {
		return cache[key]
	}

	if len(springs) == 0 {
		if strings.Contains(lava, "#") {
			return 0
		}
		return 1
	}

	current := springs[0]
	springs = springs[1:]
	result := 0
	for i := 0; i < len(lava)-sum(springs)-len(springs)-current+1; i++ {
		if strings.Contains(lava[:i], "#") {
			break
		}
		nxt := i + current
		if nxt > len(lava) {
			break
		}
		if strings.Contains(lava[i:nxt], ".") {
			continue
		}
		if nxt == len(lava) {
			result += 1
		} else if lava[nxt:nxt+1] != "#" {
			result += recurse(lava[nxt+1:], springs)
		}
	}

	cache[key] = result
	return result
}

func sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func repeat(slice []int, count int) []int {
	result := make([]int, len(slice)*count)
	for i := range result {
		result[i] = slice[i%len(slice)]
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total_ways := 0
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")
		lava := data[0]
		springs := strings.Split(data[1], ",")
		springInts := make([]int, len(springs))
		for i, spring := range springs {
			springInts[i], _ = strconv.Atoi(spring)
		}
		combinations := recurse(lava, springInts)
		total_ways += combinations
	}
	fmt.Println(total_ways)
}
