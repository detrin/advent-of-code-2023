package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func matchDigit(s string) string {
	for word, digit := range digitMap {
		if strings.HasPrefix(s, word) {
			return digit
		}
	}
	return ""
}

func findDigits(s string) (string, string) {
	firstDigit, lastDigit := "", ""
	for pos := 0; pos < len(s); pos++ {
		ch := rune(s[pos])
		if unicode.IsLetter(ch) {
			digit := matchDigit(s[pos:])
			if digit != "" {
				if firstDigit == "" {
					firstDigit = digit
				}
				lastDigit = digit
			}
		} else if unicode.IsDigit(ch) {
			if firstDigit == "" {
				firstDigit = string(ch)
			}
			lastDigit = string(ch)
		}
	}
	return firstDigit, lastDigit
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := findDigits(strings.ToLower(line))
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
