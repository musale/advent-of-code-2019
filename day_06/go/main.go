package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type position struct {
	previous string
	current  string
	next     string
}

func main() {
	positions := make(map[string][]position)
	mapData, err := ioutil.ReadFile("day_06/input.txt")
	if err != nil {
		log.Fatalf("failed to open file with error %v", err)
	}
	mapValues := strings.Split(string(mapData), "\n")
	for i := 0; i < len(mapValues); i++ {
		var prev, curr, next string
		mapValue := strings.Split(string(mapValues[i]), ")")
		if i == 0 {
			prev = ""
		} else {
			m := strings.Split(string(mapValues[i-1]), ")")
			prev = m[1]
		}
		curr = mapValue[0]
		if i+1 > len(mapValues) {
			next = ""
		} else {
			next = mapValue[1]
		}
		node := position{prev, curr, next}
		a := positions[curr]
		positions[curr] = append(a, node)
	}
	sum := 0
	for p := range positions {
		sum += loop(p, positions)
	}
	fmt.Println(sum)
}

// loop - adds the total of the sub trees
func loop(pos string, positions map[string][]position) int {
	sum := 0
	for _, p := range positions[pos] {
		sum += loop(p.next, positions)
		sum++
	}
	return sum
}
