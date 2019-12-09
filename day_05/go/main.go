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
	Modes    []code
	Instructions []code
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
	c.Modes = []code{(mode / 100) % 10, (mode / 1000) % 10, (mode / 10000) % 10}
	switch opcode {
	case 1:
		c.Input[c.getAddr(3)] = c.getArgument(1) + c.getArgument(2)
		c.StartPos += 4
	case 2:
		c.Input[c.getAddr(3)] = c.getArgument(1) * c.getArgument(2)
		c.StartPos += 4
	case 3:
		if len(c.Instructions) > 0{
			c.Input[c.getAddr(1)] = c.Instructions[0]
			c.Instructions = c.Instructions[1:]
			c.StartPos += 2
		} else{
			c.Paused = true
		}
	case 4:
		c.Outputs = append(c.Outputs, c.getArgument(1))
		c.StartPos +=2
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
	inputValues, err := ioutil.ReadFile("day_05/input.txt")
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

// maxOutput - gets the largest value in the output array
func (c *IntCode) maxOutput() code {
	var max code
	for i := 0; i < len(c.Outputs); i++ {
		curr := c.Outputs[i]
		if curr > max{
			max = curr
		}
	}
	return max
}

func main() {
	partOneInput := readInput()
	var intCode IntCode
	intCode.Running = true
	intCode.Input = partOneInput
	intCode.Instructions = []code{code(1)}
	intCode.Run()
	fmt.Println("Part 1:", intCode.maxOutput())
}
