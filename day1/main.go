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

	for i, v := range input {
		if i >= 1 && v > input[i-1] {
			sum = sum + 1
		}
	}

	fmt.Printf("Increases: %d", sum)
}
