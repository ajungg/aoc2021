package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var origLines []string
var lines []string

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines = make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	origLines = lines

	oxygenRating := getValue(0)
	co2Rating := getValue(1)

	fmt.Println(oxygenRating * co2Rating)
}

func getValue(offset byte) int {
	result := 0
	bitSize := len(lines[0])
	lines = origLines

	for i := 0; i < bitSize; i++ {
		bit := getBit(i, offset)
		p := bitSize - i - 1
		if bit == '1' {
			result = result | 1<<p
		} else {
			result = result | 0<<p
		}
		keepValuesWithBitsAtIndex(i, bit)
	}
	return result
}

func getBit(index int, offset byte) rune {
	ones := 0
	zeros := 0

	for _, line := range lines {
		if len(lines) == 1 {
			return rune(line[index])
		}

		if line[index] == '1' {
			ones += 1
		} else {
			zeros += 1
		}
	}

	if ones >= zeros {
		return rune('1' - offset)
	} else {
		return rune('0' + offset)
	}
}

func keepValuesWithBitsAtIndex(index int, bit rune) {
	newLines := make([]string, 0)
	for _, v := range lines {
		if v[index] == byte(bit) {
			newLines = append(newLines, v)
		}
	}
	lines = newLines
}
