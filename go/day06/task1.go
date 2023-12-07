package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getWaysToWin(time, record int) int {
	a := -1.0
	b := float64(time)
	c := -float64(record)
	discriminant := math.Sqrt(b*b - 4*a*c)
	if discriminant < 0 {
		return 0
	}

	lower := (-b + discriminant) / (2.0 * a)
	upper := (-b - discriminant) / (2.0 * a)
	if math.Ceil(lower) != lower {
		lower = math.Ceil(lower)
	} else {
		lower += 1
	}
	if math.Floor(upper) != upper {
		upper = math.Floor(upper)
	} else {
		upper -= 1
	}

	numWays := int(upper - lower + 1)
	if numWays > 0 {
		return numWays
	} else {
		return 0
	}
}

func readInts(scanner *bufio.Scanner, prefix string) []int {
	scanner.Scan()
	line := strings.TrimPrefix(scanner.Text(), prefix)
	fields := strings.Fields(line)
	ints := make([]int, len(fields))
	for i, str := range fields {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Invalid input:", str)
			os.Exit(1)
		}
		ints[i] = num
	}
	return ints
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	times := readInts(scanner, "Time:")
	records := readInts(scanner, "Distance:")

	totalWays := 1
	for i := 0; i < len(times); i++ {
		ways := getWaysToWin(times[i], records[i])
		totalWays *= ways
	}

	fmt.Println(totalWays)
}
