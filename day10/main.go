package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	now := time.Now()

	mappings := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
		'<': '>',
	}

	scorings := map[rune]int{
		')': 3,
		'}': 57,
		']': 1197,
		'>': 25137,
	}

	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		a := make([]rune, 0)
		for _, r := range line {
			if _, ok := mappings[r]; ok {
				a = append(a, r)
			} else {
				last := a[len(a)-1]

				if mappings[last] == r {
					a = a[:len(a)-1]
				} else {
					score += scorings[r]
					break
				}
			}
		}
	}

	fmt.Printf("time = %fs\n", time.Since(now).Seconds())
	fmt.Printf("result = %d", score)
}

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
