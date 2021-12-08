package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	values := strings.Split(string(input), ",")

	positions := make([]int, len(values))
	for i, v := range values {
		value, _ := strconv.Atoi(v)
		positions[i] = value
	}

	now := time.Now()
	fuelNeeded := make(map[int]int)

	lookupTable := make([]int, 2000)
	for i := 0; i < len(lookupTable); i++ {
		lookupTable[i] = i * (i + 1) / 2
	}

	for i := 0; i < len(positions); i++ {
		for _, v := range positions {
			fuelNeeded[i] += lookupTable[abs(v-i)]
		}
	}

	minFuelNeeded := math.MaxInt
	for _, v := range fuelNeeded {
		if v < minFuelNeeded {
			minFuelNeeded = v
		}
	}

	fmt.Printf("time = %fs\n", time.Since(now).Seconds())
	fmt.Printf("result = %d\n", minFuelNeeded)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
