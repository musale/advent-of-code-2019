package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	start := 183564
	stop := 657474
	counts := generatePasswords(start, stop)
	fmt.Println(counts)
}

func meetsCriteria(num int) bool {
	adjacent := false
	increases := false
	trip := false
	numStr := strconv.Itoa(num)
	numStrVals := strings.Split(numStr, "")
	adjCounts, incCounts, tripCounts := 0, 0, 0
	for i := 0; i < len(numStrVals)-1; i++ {
		curr, _ := strconv.Atoi(string(numStr[i]))
		next, _ := strconv.Atoi(string(numStr[i+1]))
		if next >= curr {
			incCounts++
		}
		if curr == next {
			adjCounts++
		}
	}
	for i := 0; i < len(numStrVals)-3; i++ {
		curr, _ := strconv.Atoi(string(numStr[i]))
		next, _ := strconv.Atoi(string(numStr[i+1]))
		nextOne, _ := strconv.Atoi(string(numStr[i+2]))
		nextTwo, _ := strconv.Atoi(string(numStr[i+3]))
		if curr == next && next == nextOne {
			if nextOne != nextTwo {
				tripCounts++
			}
		}
	}

	if incCounts == 5 {
		increases = true
	}
	if adjCounts >= 1 {
		adjacent = true
	}
	if tripCounts > 0 {
		trip = true
	}
	if increases && adjacent && !trip {
		return true
	}
	return false
}

func generatePasswords(start, stop int) int {
	counts := 0
	for start != stop {
		start++
		if meetsCriteria(start) {
			counts++
		}
	}
	return counts
}
