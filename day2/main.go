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

	x := 0
	z := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		direction := split[0]
		distance, _ := strconv.Atoi(split[1])

		if direction == "forward" {
			x += distance
		} else if direction == "up" {
			z -= distance
		} else if direction == "down" {
			z += distance
		}
	}

	fmt.Printf("%d\n", x*z)
}
