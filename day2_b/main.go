package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	horizontal := 0
	depth := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		direction := split[0]
		distance, _ := strconv.Atoi(split[1])

		if direction == "forward" {
			horizontal += distance
			depth += aim * distance
		} else if direction == "up" {
			aim -= distance
		} else if direction == "down" {
			aim += distance
		}
	}

	fmt.Printf("%d\n", horizontal*depth)
}
