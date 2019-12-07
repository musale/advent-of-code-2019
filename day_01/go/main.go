package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {

		value := scanner.Text()
		v, _ := strconv.Atoi(value)
		total += getMass(v, 0)
		// total += mass
	}
	fmt.Println(total)
}

func getMass(init, total int) int {
	if init <= 0 {
		return total
	}
	newMass := (init / 3) - 2
	if newMass > 0 {
		total += newMass
	}
	return getMass(newMass, total)
}
