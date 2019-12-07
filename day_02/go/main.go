package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			stateValues, err := ioutil.ReadFile("day_02/input.txt")
			if err != nil {
				log.Fatalf("failed to open file with error %v", err)
			}
			stateStrings := strings.Split(string(stateValues), ",")
			states := make([]int, len(stateStrings))

			for j, s := range stateStrings {
				states[j], _ = strconv.Atoi(s)
			}
			i := 0
			states[1] = noun // noun
			states[2] = verb // verb
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
			if states[0] == 19690720 {
				fmt.Println(noun)
				fmt.Println(verb)
				fmt.Println(100*noun + verb)
				return
			}
		}
	}
}
