package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		value := scanner.Text()
		v, _ := strconv.Atoi(value)
		mass := (v / 3) - 2
		total += mass
	}
	fmt.Println(total)
}
