package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var octopusMap = make([]int, 0)
var w int = 0
var h int = 0
var flashes int = 0

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		if w == 0 {
			w = len(s)
		}

		for _, r := range s {
			n, _ := strconv.Atoi(string(r))
			octopusMap = append(octopusMap, n)
		}

		h += 1
	}

	for i := 0; i < 195; i++ {
		step()
	}

	fmt.Printf("%d", flashes)
}

func step() {
	increaseEnergyLevel()
	flashLoop()
}

func increaseEnergyLevel() {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			set(x, y, get(x, y)+1)
		}
	}
}

func flashLoop() {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			flash(x, y)
		}
	}
}

func flash(x int, y int) {
	v := get(x, y)
	if v > 9 {
		flashes += 1
		set(x, y, 0)

		increase(x-1, y-1)
		increase(x-1, y)
		increase(x-1, y+1)

		increase(x, y-1)
		increase(x, y+1)

		increase(x+1, y-1)
		increase(x+1, y)
		increase(x+1, y+1)
	}
}

func get(x int, y int) int {
	if x >= 0 && x < w && y >= 0 && y < h {
		return octopusMap[w*y+x]
	}
	return -1
}

func set(x int, y int, b int) {
	if x >= 0 && x < w && y >= 0 && y < h {
		octopusMap[w*y+x] = b
	}
}

func increase(x int, y int) {
	v := get(x, y)
	if v == 0 {
		return
	}

	set(x, y, v+1)
	if v+1 > 9 {
		flash(x, y)
	}
}

func debugPrint(step int) {
	fmt.Printf("Step %d:\n", step)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print(get(x, y))
		}
		fmt.Println()
	}
}
