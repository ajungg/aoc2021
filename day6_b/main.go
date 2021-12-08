package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var fishCounter map[uint64]uint64

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	values := strings.Split(string(input), ",")

	fishCounter = make(map[uint64]uint64)
	for _, v := range values {
		value, _ := strconv.Atoi(v)
		fishCounter[uint64(value)] += 1
	}

	fmt.Printf("Input values: %d\n", len(values))

	for i := 0; i < 256; i++ {
		evolve()
	}

	sum := uint64(0)
	for _, v := range fishCounter {
		sum += v
	}

	fmt.Printf("Number of fishes: %d\n", sum)
}

func evolve() {
	newFishCounter := make(map[uint64]uint64)
	for k, v := range fishCounter {
		if k == 0 {
			newFishCounter[6] += v
			newFishCounter[8] += v
		} else {
			newFishCounter[k-1] += v
		}
	}
	fishCounter = newFishCounter
}
