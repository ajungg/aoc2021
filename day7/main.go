package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	values := strings.Split(string(input), ",")

	positions := make([]int, len(values))
	for i, v := range values {
		value, _ := strconv.Atoi(v)
		positions[i] = value
	}

	fuelNeeded := make(map[int]int)
	for i, v := range positions {
		for _, v2 := range positions {
			fuelNeeded[i] += abs(v - v2)
		}
	}

	minFuelNeeded := math.MaxInt
	for _, v := range fuelNeeded {
		if v < minFuelNeeded {
			minFuelNeeded = v
		}

	}

	fmt.Printf("%d\n", minFuelNeeded)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
