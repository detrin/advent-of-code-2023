package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	total := 0
	for _, step := range steps {
		total += hash(step)
	}

	fmt.Println(total)
}
