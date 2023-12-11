package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func getNext(arr []int) int {
    n := len(arr)
	all_zero := true
    diffs := make([]int, n-1)
    for i := 0; i < n-1; i++ {
        diffs[i] = arr[i+1] - arr[i]
		if diffs[i] != 0 {
			all_zero = false
		}
    }

    if !all_zero {
		return arr[n-1] + getNext(diffs)
	} 

	return arr[n-1]
}

func main() {
    histories := make([][]int, 0)
	scanner := bufio.NewScanner(os.Stdin)    
	for scanner.Scan() {  
		line := scanner.Text()
		nums := strings.Split(line, " ")
		history := make([]int, len(nums))
		for i, num := range nums {
			history[i], _ = strconv.Atoi(num)
		}
		histories = append(histories, history)
	}

    sum := 0
    for _, history := range histories {
        next := getNext(history)
        sum += next
    }

    fmt.Println(sum)
}
