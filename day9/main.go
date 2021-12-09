package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	riskLevel := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			v := get(x, y)

			top := get(x, y-1)
			left := get(x-1, y)
			right := get(x+1, y)
			bottom := get(x, y+1)

			if top > v && left > v && right > v && bottom > v {
				riskLevel += 1 + int(v)
			}
		}
	}

	fmt.Printf("time = %fs\n", time.Since(now).Seconds())
	fmt.Printf("result = %d\n", riskLevel)
}

func get(x int, y int) byte {
	if x >= 0 && x < w && y >= 0 && y < h {
		return heightmap[w*y+x]
	}
	return 9
}
