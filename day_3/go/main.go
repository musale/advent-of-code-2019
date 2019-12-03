package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

// Point - is an x,y coordinate with float64 x & y axis values
type Point struct {
	x float64
	y float64
}

// contains - checks if a Point exists in a map of Points
func contains(points map[Point]int, point Point) bool {
	_, ok := points[point]
	return ok
}

func main() {
	distanceValues, err := ioutil.ReadFile("day_3/input.txt")
	if err != nil {
		log.Fatalf("failed to open file with error %v", err)
	}
	distanceStrings := strings.Split(string(distanceValues), "\n")
	wireOneStrings := distanceStrings[0]
	wireTwoStrings := distanceStrings[1]

	wireOneValues := strings.Split(wireOneStrings, ",")
	wireTwoValues := strings.Split(wireTwoStrings, ",")

	wireOnePoints := getPoints(wireOneValues)
	wireTwoPoints := getPoints(wireTwoValues)

	fmt.Println(len(wireOnePoints))
	fmt.Println(len(wireTwoPoints))

	commonPoints := []Point{}
	for point := range wireOnePoints {
		if contains(wireTwoPoints, point) {
			commonPoints = append(commonPoints, point)
		}
	}
	totals := []float64{}
	for _, point := range commonPoints {
		totals = append(totals, point.x+point.y)
	}
	fmt.Println(totals)
	min, max := minMax(totals)
	fmt.Println(min)
	fmt.Println(max)
}

func getPoints(arr []string) map[Point]int {
	XD := map[string]int{"R": 1, "L": -1, "U": 0, "D": 0}
	YD := map[string]int{"R": 0, "L": 0, "U": 1, "D": -1}

	x, y, steps := 0, 0, 0
	ans := make(map[Point]int)
	for _, s := range arr {
		direction := string(s[0])
		distance, _ := strconv.Atoi(s[1:])
		for distance > 0 {
			steps++
			x += XD[direction]
			y += YD[direction]
			a := Point{math.Abs(float64(x)), math.Abs(float64(y))}
			if !contains(ans, a) {
				ans[a] = steps
			}
			distance--
		}
	}
	return ans
}

func minMax(array []float64) (float64, float64) {
	var max float64 = array[0]
	var min float64 = array[0]
	for _, value := range array {
		if max <= value {
			max = value
		}
		if min >= value {
			min = value
		}
	}
	return min, max
}
