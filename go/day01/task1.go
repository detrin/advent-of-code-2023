package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := ""
		lastDigit := ""
		for _, r := range line {
			if unicode.IsDigit(r) {
				if firstDigit == "" {
					firstDigit = string(r)
				}
				lastDigit = string(r)
			}
		}
		if firstDigit != "" && lastDigit != "" {
			number, _ := strconv.Atoi(firstDigit + lastDigit)
			sum += number
		}
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", scanner.Err())
	}
	fmt.Println(sum)
}
