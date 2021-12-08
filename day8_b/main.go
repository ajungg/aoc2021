package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	result := 0
	for scanner.Scan() {
		line := string(scanner.Text())
		values := strings.Split(line, " | ")
		codeWords := strings.Split(values[0], " ")
		outputWords := strings.Split(values[1], " ")

		counter := make(map[string]int)
		for _, cw := range codeWords {
			counter[cw] += 1
		}

		// 1, 4, 7 and 8 are unique
		// 0 = 6 = 9 --> 6 bits
		// 2 = 3 = 5 --> 5 bits

		bitmap := make(map[int]string)

		for k, v := range counter {
			if v == 1 {
				switch len(k) {
				case 2:
					bitmap[1] = k
				case 3:
					bitmap[7] = k
				case 4:
					bitmap[4] = k
				case 7:
					bitmap[8] = k
				}
			}
		}

		for k, _ := range counter {
			if len(k) == 6 && ContainsAllRunes(k, bitmap[4]) && ContainsAllRunes(k, bitmap[7]) {
				bitmap[9] = k
				break
			}
		}

		for k, _ := range counter {
			if len(k) == 6 && ContainsAllRunes(k, bitmap[1]) && ContainsAllRunes(k, bitmap[7]) && k != bitmap[9] {
				bitmap[0] = k
			} else if len(k) == 6 && !ContainsAllRunes(k, bitmap[1]) {
				bitmap[6] = k
			}
		}

		for k, _ := range counter {
			if len(k) == 5 && ContainsAllRunes(k, bitmap[1]) {
				bitmap[3] = k
				break
			}
		}

		bitC := ReplaceAllRunes(bitmap[1], bitmap[6])

		for k, _ := range counter {
			if (len(k) == 5) && !ContainsAllRunes(k, bitC) {
				bitmap[5] = k
				break
			}
		}

		for k, _ := range counter {
			if (len(k) == 5) && k != bitmap[3] && k != bitmap[5] {
				bitmap[2] = k
				break
			}
		}

		reverseBitmap := make(map[string]int)
		for k, v := range bitmap {
			reverseBitmap[v] = k
		}
		sb := strings.Builder{}

		for _, ow := range outputWords {
			for k, v := range reverseBitmap {
				if len(ow) == len(k) && ContainsAllRunes(ow, k) {
					s := strconv.Itoa(v)
					sb.WriteString(s)
				}
			}
		}

		value, _ := strconv.Atoi(sb.String())
		result += value
	}

	fmt.Printf("time = %fs\n", time.Since(now).Seconds())
	fmt.Printf("result = %d\n", result)
}

func ReplaceAllRunes(input string, replace string) string {
	output := input
	for _, r := range replace {
		output = strings.ReplaceAll(output, string(r), "")
	}
	return output
}

func ContainsAllRunes(input string, all string) bool {
	for _, r := range all {
		if !strings.ContainsRune(input, r) {
			return false
		}
	}
	return true
}
