package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	stateValues, err := ioutil.ReadFile("day_2/input.txt")
	if err != nil {
		log.Fatalf("failed to open file with error %v", err)
	}
	stateStrings := strings.Split(string(stateValues), ",")
	states := make([]int, len(stateStrings))

	for j, s := range stateStrings {
		states[j], _ = strconv.Atoi(s)
	}

	i := 0
	states[1] = 12
	states[2] = 2
	for states[i] != 99 {
		switch states[i] {
		case 1:
			// add
			sum := states[states[i+1]] + states[states[i+2]]
			states[states[i+3]] = sum
			i += 4
		case 2:
			// multiply
			mul := states[states[i+1]] * states[states[i+2]]
			states[states[i+3]] = mul
			i += 4
		}
	}
	fmt.Println(states)
}
