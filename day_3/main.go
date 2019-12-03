package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"reflect"
	"strconv"
	"strings"
)

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

	commonPoints := [][2]float64{}
	wireOneKeys := reflect.ValueOf(wireTwoPoints).MapKeys()
	for _, p1 := range wireOneKeys {
		point := [2]float64{p1.Index(0).Float(), p1.Index(1).Float()}
		_, ok := wireOnePoints[point]
		if ok {
			commonPoints = append(commonPoints, point)
		}
	}
	totals := []float64{}
	for _, point := range commonPoints {
		totals = append(totals, point[0]+point[1])
	}
	fmt.Println(totals)
	min, max := minMax(totals)
	fmt.Println(min)
	fmt.Println(max)
}

func getPoints(arr []string) map[[2]float64]int {
	XD := map[string]int{"R": 1, "L": -1, "U": 0, "D": 0}
	YD := map[string]int{"R": 0, "L": 0, "U": 1, "D": -1}

	x, y, l := 0, 0, 0
	ans := make(map[[2]float64]int)
	for _, s := range arr {
		direction := string(s[0])
		distance, _ := strconv.Atoi(s[1:])
		for distance > 0 {
			l++
			x += XD[direction]
			y += YD[direction]
			a := [2]float64{math.Abs(float64(x)), math.Abs(float64(y))}
			_, ok := ans[a]
			if ok == false {
				ans[a] = l
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
