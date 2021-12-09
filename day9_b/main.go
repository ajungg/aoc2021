package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

var heightmap []byte = make([]byte, 0)
var w int = 0
var h int = 0

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
			heightmap = append(heightmap, byte(n))
		}

		h += 1
	}

	now := time.Now()

	basins := make([]int, 0)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			v := get(x, y)

			top := get(x, y-1)
			left := get(x-1, y)
			right := get(x+1, y)
			bottom := get(x, y+1)

			if top > v && left > v && right > v && bottom > v {
				basins = append(basins, getSizeOfBasin(x, y))
			}
		}
	}

	sort.Ints(basins)
	result := basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]

	fmt.Printf("time = %fs\n", time.Since(now).Seconds())
	fmt.Printf("result = %d", result)
}

func getSizeOfBasin(x int, y int) int {
	count := 0
	if get(x, y) < 9 {
		count += 1
		heightmap[w*y+x] = 9
		count += getSizeOfBasin(x, y+1)
		count += getSizeOfBasin(x, y-1)
		count += getSizeOfBasin(x+1, y)
		count += getSizeOfBasin(x-1, y)
	}
	return count
}

func get(x int, y int) byte {
	if x >= 0 && x < w && y >= 0 && y < h {
		return heightmap[w*y+x]
	}
	return 9
}
