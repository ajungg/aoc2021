package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var generation []int8

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	values := strings.Split(string(input), ",")
	generation = make([]int8, len(values), 100000)

	for i, v := range values {
		value, _ := strconv.Atoi(v)
		generation[i] = int8(value)
	}

	fmt.Printf("Input values: %d\n", len(generation))

	for i := 0; i < 80; i++ {
		evolve()
	}

	fmt.Printf("Number of fishes: %d\n", len(generation))
}

func evolve() {
	for i, v := range generation {
		newValue := v - 1
		if newValue < 0 {
			generation = append(generation, 8)
			generation[i] = 6
		} else {
			generation[i] = newValue
		}
	}
}

var day int = 0

func printGeneration() {
	fmt.Printf("%d   ", day)

	for _, v := range generation {
		fmt.Print(v)
	}
	fmt.Println()

	day = day + 1
}
