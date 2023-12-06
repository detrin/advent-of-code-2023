package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type SeedRange struct {
	Start int
	Range int
}

type Converter struct {
	From string
	To   string
	Map  [][]int
}

func parseSeeds(seedsLine string) []SeedRange {
	seedStrings := strings.Fields(strings.TrimPrefix(seedsLine, "seeds:"))
	if len(seedStrings)%2 != 0 {
		return nil
	}

	seeds := make([]SeedRange, len(seedStrings)/2)
	for i := 0; i < len(seedStrings); i += 2 {
		start, _ := strconv.Atoi(seedStrings[i])
		range_, _ := strconv.Atoi(seedStrings[i+1])
		seeds[i/2] = SeedRange{Start: start, Range: range_}
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

func parseInput(input string) ([]SeedRange, []Converter) {
	lines := strings.Split(input, "\n")
	seeds := parseSeeds(lines[0])
	converterStrings := strings.Split(strings.Join(lines[2:], "\n"), "\n\n")
	converters := make([]Converter, len(converterStrings))
	for i, cs := range converterStrings {
		converters[i] = parseConverter(cs)
	}
	return seeds, converters
}

func convert(seedRange SeedRange, converter Converter) []SeedRange {
	newSeedRanges := []SeedRange{}

	remainingRange := SeedRange{Start: seedRange.Start, Range: seedRange.Range}

	for _, mapping := range converter.Map {
		startDst, startSrc, length := mapping[0], mapping[1], mapping[2]
		endSrc := startSrc + length

		if remainingRange.Start < endSrc && remainingRange.Start+remainingRange.Range > startSrc {
			intersectStart := max(remainingRange.Start, startSrc)
			intersectEnd := min(remainingRange.Start+remainingRange.Range, endSrc)

			newStart := startDst + (intersectStart - startSrc)
			newEnd := startDst + (intersectEnd - startSrc)
			if newStart < newEnd {
				newSeedRanges = append(newSeedRanges, SeedRange{Start: newStart, Range: newEnd - newStart})
			}

			if remainingRange.Start == intersectStart {
				remainingRange.Start = intersectEnd
				remainingRange.Range -= intersectEnd - intersectStart
			} else if remainingRange.Start+remainingRange.Range == intersectEnd {
				remainingRange.Range = intersectStart - remainingRange.Start
			}
		}
	}

	if remainingRange.Range > 0 {
		newSeedRanges = append(newSeedRanges, remainingRange)
	}

	return newSeedRanges
}

func mergeSeedRanges(seedRanges []SeedRange) []SeedRange {
	seen := make(map[SeedRange]bool)
	uniqueSeedRanges := []SeedRange{}
	for _, seedRange := range seedRanges {
		if !seen[seedRange] {
			seen[seedRange] = true
			uniqueSeedRanges = append(uniqueSeedRanges, seedRange)
		}
	}

	sort.Slice(uniqueSeedRanges, func(i, j int) bool {
		return uniqueSeedRanges[i].Start < uniqueSeedRanges[j].Start
	})

	merged := []SeedRange{}
	for _, seedRange := range uniqueSeedRanges {
		if len(merged) == 0 || merged[len(merged)-1].Start+merged[len(merged)-1].Range <= seedRange.Start {
			merged = append(merged, seedRange)
		} else {
			merged[len(merged)-1].Range = max(merged[len(merged)-1].Range, seedRange.Start+seedRange.Range-merged[len(merged)-1].Start)
		}
	}

	return merged
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	input := strings.Join(text, "\n")
	seed_store1, converters := parseInput(input)
	seed_store2 := make([]SeedRange, 0)
	p_seed_store1 := &seed_store1
	p_seed_store2 := &seed_store2

	converterMap := make(map[string]Converter)
	for _, converter := range converters {
		converterMap[converter.From] = converter
	}

	state := "seed"
	is_converter_available := true
	for is_converter_available {
		converter, ok := converterMap[state]
		if ok {
			*p_seed_store2 = make([]SeedRange, 0)
			for _, seedRange := range *p_seed_store1 {
				seed_ranges := convert(seedRange, converter)
				*p_seed_store2 = append(*p_seed_store2, seed_ranges...)
			}
			p_seed_store1, p_seed_store2 = p_seed_store2, p_seed_store1
			*p_seed_store1 = mergeSeedRanges(*p_seed_store1)
			state = converter.To
		} else {
			is_converter_available = false
		}
	}

	minSeed := (*p_seed_store1)[0].Start
	for _, seedRange := range *p_seed_store1 {
		if seedRange.Start < minSeed {
			minSeed = seedRange.Start
		}
	}
	fmt.Println(minSeed)
}
