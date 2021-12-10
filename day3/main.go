package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var counts []int

	for scanner.Scan() {
		line := scanner.Text()

		if counts == nil {
			counts = make([]int, len(line))
		}

		for i, c := range line {
			if c == '1' {
				counts[i] += 1
			} else {
				counts[i] -= 1
			}
		}
	}

	gamma := 0
	epsilon := 0

	for i, c := range counts {
		b := len(counts) - i - 1
		if c > 0 {
			gamma = gamma | 1<<b
			epsilon = epsilon | 0<<b
		} else {
			gamma = gamma | 0<<b
			epsilon = epsilon | 1<<b
		}
	}

	fmt.Println(gamma * epsilon)
}
