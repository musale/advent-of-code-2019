package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// code - single integer in an IntCode program
type code int

// IntCode - program for a machine that can run int code
type IntCode struct {
	Input    []code
	Outputs  []code
	Modes    [3]code
	StartPos code
	RelBase  code
	Running  bool
	Paused   bool
}

// Tick - starts an int code program with the provided input
func (c *IntCode) Tick() {
	// Calculate the modes
	mode := c.Input[c.StartPos]
	opcode := mode % 100
	c.Modes = [3]code{(mode / 100) % 10, (mode / 1000) % 10, (mode / 10000) % 10}
	c.Paused = false
	switch opcode {
	case 1:
		c.Input[c.getAddr(3)] = c.getArgument(1) + c.getArgument(2)
		c.StartPos += 4
	case 2:
		c.Input[c.getAddr(3)] = c.getArgument(1) * c.getArgument(2)
		c.StartPos += 4
	case 99:
		c.Running = false
	}
}

// Run - an IntCode machine
func (c *IntCode) Run() {
	for c.Running != false {
		c.Tick()
	}
}

// getArgument - of a given position
func (c *IntCode) getArgument(pos code) code {
	arg := c.Modes[pos-1]
	var val code
	switch arg {
	case 0:
		val = c.Input[c.Input[c.StartPos+pos]]
	case 1:
		val = c.Input[c.StartPos+pos]
	case 2:
		val = c.Input[c.RelBase+c.Input[c.StartPos+pos]]
	default:
		log.Fatalf("Unknown argument mode %d", arg)
	}
	return val
}

// getAddr - at a given position
func (c *IntCode) getAddr(pos code) code {
	arg := c.Modes[pos-1]
	var val code
	switch arg {
	case 0:
		val = c.Input[c.StartPos+pos]
	case 1:
		log.Fatalf("Failed getting arg at %d with immediate argument mode", arg)
	case 2:
		val = c.RelBase + c.Input[c.StartPos+pos]
	default:
		log.Fatalf("Unknown argument mode %d", arg)
	}
	return val
}

// readInput - from file and convert to code
func readInput() []code {
	inputValues, err := ioutil.ReadFile("day_02/input.txt")
	if err != nil {
		log.Fatalf("Failed to open file with error %v", err)
	}
	stateStrings := strings.Split(string(inputValues), ",")
	input := make([]code, len(stateStrings))

	for j, s := range stateStrings {
		num, _ := strconv.Atoi(s)
		input[j] = code(num)
	}
	return input
}

func main() {
	partOneInput := readInput()
	var intCode IntCode
	intCode.Running = true
	intCode.Input = partOneInput
	intCode.Run()
	fmt.Println("Part 1:", intCode.Input[0])

	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			partTwoInput := readInput()
			partTwoInput[1] = code(noun)
			partTwoInput[2] = code(verb)
			var intCode IntCode
			intCode.Input = partTwoInput
			intCode.Running = true
			intCode.Run()
			if intCode.Input[0] == 19690720 {
				fmt.Println(noun)
				fmt.Println(verb)
				fmt.Println("Part 2:", 100*noun+verb)
				return
			}
		}
	}
}
