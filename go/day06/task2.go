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

func readInt(scanner *bufio.Scanner, prefix string) int {
	scanner.Scan()
	line := strings.TrimPrefix(scanner.Text(), prefix)
	line = strings.ReplaceAll(line, " ", "")
	num, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("Invalid input:", line)
		os.Exit(1)
	}
	return num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	time := readInt(scanner, "Time:")
	record := readInt(scanner, "Distance:")

	ways := getWaysToWin(time, record)

	fmt.Println(ways)
}
