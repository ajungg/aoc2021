package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	now := time.Now()
	scanner := bufio.NewScanner(file)

	easyWordCount := 0
	for scanner.Scan() {
		line := string(scanner.Text())
		values := strings.Split(line, " | ")
		codeWords := strings.Split(values[0], " ")
		outputWords := strings.Split(values[1], " ")

		counter := make(map[int]int)
		for _, cw := range codeWords {
			counter[len(cw)] += 1
		}

		for _, cw := range outputWords {
			if counter[len(cw)] == 1 {
				easyWordCount += 1
			}
		}
	}

	fmt.Printf("time = %fs\n", time.Since(now).Seconds())
	fmt.Printf("result = %d\n", easyWordCount)
}
