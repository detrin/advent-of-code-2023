package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Converter struct {
	From string
	To   string
	Map  [][]int
}

func parseSeeds(seedsLine string) []int {
	seedStrings := strings.Fields(strings.TrimPrefix(seedsLine, "seeds:"))
	seeds := make([]int, len(seedStrings))
	for i, s := range seedStrings {
		seeds[i], _ = strconv.Atoi(s)
	}
	return seeds
}

func parseConverter(input string) Converter {
	lines := strings.Split(input, "\n")
	fromTo := strings.Split(lines[0], "-to-")
	from, to := strings.TrimSpace(fromTo[0]), strings.Fields(fromTo[1])[0]

	maps := make([][]int, len(lines)-1)
	for i, line := range lines[1:] {
		nums := strings.Fields(line)
		maps[i] = make([]int, len(nums))
		for j, num := range nums {
			maps[i][j], _ = strconv.Atoi(num)
		}
	}

	return Converter{
		From: from,
		To:   to,
		Map:  maps,
	}
}

func parseInput(input string) ([]int, []Converter) {
	lines := strings.Split(input, "\n")
	seeds := parseSeeds(lines[0])
	converterStrings := strings.Split(strings.Join(lines[2:], "\n"), "\n\n")
	converters := make([]Converter, len(converterStrings))
	for i, cs := range converterStrings {
		converters[i] = parseConverter(cs)
	}
	return seeds, converters
}

func convert(n int, converter Converter) int {
	for _, mapping := range converter.Map {
		startDst, startSrc, length := mapping[0], mapping[1], mapping[2]
		if n >= startSrc && n < startSrc+length {
			return startDst + (n - startSrc)
		}
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	input := strings.Join(text, "\n")
	seeds_store1, converters := parseInput(input)
	seeds_store2 := make([]int, 0)

	p_seeds_store1 := &seeds_store1
	p_seeds_store2 := &seeds_store2

	converterMap := make(map[string]Converter)
	for _, converter := range converters {
		converterMap[converter.From] = converter
	}

	state := "seed"
	is_converter_available := true
	for is_converter_available {
		converter, ok := converterMap[state]
		if ok {
			*p_seeds_store2 = make([]int, 0)
			for _, seed := range *p_seeds_store1 {
				seed = convert(seed, converter)
				*p_seeds_store2 = append(*p_seeds_store2, seed)
			}
			p_seeds_store1, p_seeds_store2 = p_seeds_store2, p_seeds_store1
			state = converter.To
		} else {
			is_converter_available = false
		}
	}

	seed_min := (*p_seeds_store1)[0]
	for _, seed := range *p_seeds_store1 {
		if seed < seed_min {
			seed_min = seed
		}
	}
	fmt.Println(seed_min)
}
