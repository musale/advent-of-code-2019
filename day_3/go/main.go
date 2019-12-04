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
	commonValues := pointIntersection(wireOnePoints, wireTwoPoints)
	fmt.Println(len(wireOnePoints))
	totals := []float64{}
	for _, point := range commonValues {
		t := point.x + point.y
		totals = append(totals, t)
	}

	min := totals[0]
	for _, t := range totals {
		if t < min {
			min = t
		}
	}
	print(min)
}

func minPoint(points []Point) Point {
	p := points[0]
	for _, point := range points {
		if point.x < p.x && point.y < p.y {
			p = point
		} else {
			if point.y < p.y {
				p = point
			}
		}
	}
	return p
}

func pointIntersection(w1, w2 map[Point]int) []Point {
	points := []Point{}
	for point := range w1 {
		_, ok := w2[point]
		if ok {
			points = append(points, point)
		}
	}
	return points
}

func getPoints(arr []string) map[Point]int {
	XD := map[string]int{"R": 1, "L": -1, "U": 0, "D": 0}
	YD := map[string]int{"R": 0, "L": 0, "U": 1, "D": -1}

	x, y := 0, 0
	ans := make(map[Point]int)
	for _, s := range arr {
		steps := 0
		direction := string(s[0])
		distance, _ := strconv.Atoi(s[1:])
		for i := 0; i < distance; i++ {
			steps++
			x += XD[direction]
			y += YD[direction]
			a := Point{math.Abs(float64(x)), math.Abs(float64(y))}
			ans[a] = steps
		}
	}
	return ans
}
