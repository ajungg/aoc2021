package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	input := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		input = append(input, n)
	}

	lastSlidingWindow := 0
	for i := range input {
		if i <= 2 {
			continue
		}

		slidingWindow := input[i] + input[i-1] + input[i-2]

		if slidingWindow > lastSlidingWindow {
			sum = sum + 1
		}

		lastSlidingWindow = slidingWindow
	}

	fmt.Printf("Increases: %d\n", sum)
}
