package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	scores := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		a := make([]rune, 0)
		broken := false
		for _, r := range line {
			if _, ok := mappings[r]; ok {
				a = append(a, r)
			} else {
				last := a[len(a)-1]

				if mappings[last] == r {
					a = a[:len(a)-1]
				} else {
					broken = true
					break
				}
			}
		}

		if !broken {
			score := 0
			for i := len(a) - 1; i >= 0; i-- {
				score = score*5 + points[mappings[a[i]]]
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	fmt.Printf("time = %fs\n", time.Since(now).Seconds())
	fmt.Printf("result = %d", scores[len(scores)/2])
}

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
