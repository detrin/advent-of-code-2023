package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	label string
	power int
}

func hash(s string) int {
	current := 0
	for _, ch := range s {
		current += int(ch)
		current *= 17
		current %= 256
	}
	return current
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	steps := strings.Split(line, ",")
	boxes := make([][]Lens, 256)

	for _, step := range steps {
		step = strings.Replace(step, "-", "", -1)
		parts := strings.Split(step, "=")
		label := parts[0]
		boxIndex := hash(label)
		if len(parts) == 1 {
			newBox := make([]Lens, 0)
			for _, lens := range boxes[boxIndex] {
				if lens.label != label {
					newBox = append(newBox, lens)
				}
			}
			boxes[boxIndex] = newBox
		} else {
			power, _ := strconv.Atoi(parts[1])
			newLens := Lens{label, power}

			found := false
			for i, lens := range boxes[boxIndex] {
				if lens.label == label {
					boxes[boxIndex][i] = newLens
					found = true
					break
				}
			}
			if !found {
				boxes[boxIndex] = append(boxes[boxIndex], newLens)
			}
		}
	}

	totalPower := 0
	for i, box := range boxes {
		for j, lens := range box {
			totalPower += (i + 1) * (j + 1) * lens.power
		}
	}

	fmt.Println(totalPower)
}
