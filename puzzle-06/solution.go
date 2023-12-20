package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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

type Race struct {
	time   int
	record int
}

func (r *Race) countPossibleWays() int {
	count := 0
	for hold, travel := 0, r.time; hold < r.time; {
		if hold*travel > r.record {
			count++
		}
		hold++
		travel--
	}
	return count
}

var ErrInvalidInput = errors.New("invalid input")

func filterEmpty(strs []string) []string {
	res := make([]string, 0, len(strs))
	for _, str := range strs {
		if str != "" {
			res = append(res, str)
		}
	}
	return res
}

func linesToRace(lines []string) (*Race, error) {
	timeStr := strings.Join(filterEmpty(strings.Split(strings.Split(lines[0], ":")[1], " ")), "")
	recordStr := strings.Join(filterEmpty(strings.Split(strings.Split(lines[1], ":")[1], " ")), "")
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return nil, ErrInvalidInput
	}
	record, err := strconv.Atoi(recordStr)
	if err != nil {
		return nil, ErrInvalidInput
	}
	return &Race{time, record}, nil
}

func linesToRaces(lines []string) ([]Race, error) {
	times := filterEmpty(strings.Split(strings.Split(lines[0], ":")[1], " "))
	records := filterEmpty(strings.Split(strings.Split(lines[1], ":")[1], " "))
	if len(times) != len(records) {
		return nil, ErrInvalidInput
	}
	races := make([]Race, 0, len(times))
	for i := range times {
		t, err := strconv.Atoi(times[i])
		if err != nil {
			return nil, ErrInvalidInput
		}
		r, err := strconv.Atoi(records[i])
		if err != nil {
			return nil, ErrInvalidInput
		}
		races = append(races, Race{
			time:   t,
			record: r,
		})
	}
	return races, nil
}

func getTotalWays(races []Race) int {
	total := 1
	for _, race := range races {
		total *= race.countPossibleWays()
	}
	return total
}

func main() {
	lines := fileToLines("input.txt")

	races, err := linesToRaces(lines)
	if err != nil {
		log.Fatalln(err.Error())
	}
	part1 := getTotalWays(races)

	race, err := linesToRace(lines)
	if err != nil {
		log.Fatalln(err.Error())
	}
	part2 := race.countPossibleWays()

	fmt.Printf("part 1: %d \npart 2: %d\n", part1, part2)
}
