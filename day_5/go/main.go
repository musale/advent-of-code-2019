package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	stateValues, err := ioutil.ReadFile("day_5/input.txt")
	if err != nil {
		log.Fatalf("failed to open file with error %v", err)
	}
	stateStrings := strings.Split(string(stateValues), ",")
	states := make([]int, len(stateStrings))

	for j, s := range stateStrings {
		states[j], _ = strconv.Atoi(s)
	}

	pos := 0
	for states[pos] != 99 {
		instructions := strings.Split(strconv.Itoa(states[pos]), "")
		modes := []int{}

		for _, m := range instructions {
			mode, _ := strconv.Atoi(m)
			modes = append(modes, mode)
		}

		opscode := modes[len(modes)-1]
		if len(modes) == 1 {
			modes = []int{0}
		} else {
			modes = modes[:len(modes)-2]
		}
		for len(modes) < 3 {
			modes = append([]int{0}, modes...)
		}
		switch opscode {
		case 1:
			for len(modes) < 3 {
				modes = append(modes, 0)
			}
			ins1, ins2, ins3 := states[pos+1], states[pos+2], states[pos+3]
			var first, second int
			if modes[2] == 1 {
				first = ins1
			} else {
				first = states[ins1]
			}
			if modes[1] == 1 {
				second = ins2
			} else {
				second = states[ins2]
			}
			states[ins3] = first + second
			pos += 4
		case 2:

			ins1, ins2, ins3 := states[pos+1], states[pos+2], states[pos+3]
			var first, second int

			if modes[2] == 1 {
				first = ins1
			} else {
				first = states[ins1]
			}
			if modes[1] == 1 {
				second = ins2
			} else {
				second = states[ins2]
			}
			states[ins3] = first * second
			pos += 4
		case 3:
			ins1 := states[pos+1]
			states[ins1] = 1 // Temporary input
			pos += 2
		case 4:
			ins1 := states[pos+1]
			fmt.Println(states[ins1])
			pos += 2
		default:
			num := states[pos]
			log.Fatalf("unknown opscode %d from state value %v", opscode, num)
		}
	}
}
