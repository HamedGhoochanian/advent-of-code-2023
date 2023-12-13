package main

import (
	"fmt"
	"os"
	"strings"
)

func fileToLines(addr string) []string {
	content, err := os.ReadFile(addr)
	if err != nil {
		return nil
	}

	lines := strings.Split(string(content), "\n")

	var result []string
	for _, line := range lines {
		if line != "" {
			result = append(result, line)
		}
	}
	return lines
}

var wordMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"0":     0,
}

type Chunk struct {
	five  string
	four  string
	three string
	digit string
}

func slidingWindow(line string) int {
	numbers := make([]int, 0)
	line = "   " + line + "     "
	length := len(line)
	for i := 0; i <= length-5; i++ {
		chunk := Chunk{
			five:  line[i : i+5],
			four:  line[i : i+4],
			three: line[i : i+3],
			digit: line[i : i+1],
		}
		if res, ok := wordMap[chunk.five]; ok {
			numbers = append(numbers, res)
		}
		if res, ok := wordMap[chunk.four]; ok {
			numbers = append(numbers, res)
		}
		if res, ok := wordMap[chunk.three]; ok {
			numbers = append(numbers, res)
		}
		if res, ok := wordMap[chunk.digit]; ok {
			numbers = append(numbers, res)
		}
	}
	first := numbers[0]
	last := numbers[len(numbers)-1]
	return first*10 + last
}


func main() {
	lines := fileToLines("input.txt")
	numbers := make([]int, 0, len(lines))
	for _, line := range lines {
		n := slidingWindow(line)
		numbers = append(numbers, n)
	}

	sum := int(0)
	for _, n := range numbers {
		sum += n
	}
	fmt.Println(sum)
}
