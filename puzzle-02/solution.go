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

type Subset struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	ID      int
	Subsets []Subset
}

func (g *Game) isPossible() bool {
	for _, set := range g.Subsets {
		if set.Red > 12 || set.Green > 13 || set.Blue > 14 {
			return false
		}
	}
	return true
}

func (g *Game) getPower() int {
	maxRed := 1
	maxGreen := 1
	maxBlue := 1
	for _, set := range g.Subsets {
		if set.Red > maxRed {
			maxRed = set.Red
		}
		if set.Blue > maxBlue {
			maxBlue = set.Blue
		}
		if set.Green > maxGreen {
			maxGreen = set.Green
		}
	}
	return maxGreen * maxBlue * maxRed
}

func setStrToSubset(setStr string) Subset {
	setStr = strings.TrimSpace(setStr)
	chunks := strings.Split(setStr, ",")
	subset := Subset{}
	for _, val := range chunks {
		trimmed := strings.TrimSpace(val)
		var count int
		var color string
		_, _ = fmt.Sscanf(trimmed, "%d %s", &count, &color)
		switch color {
		case "blue":
			subset.Blue += count
		case "red":
			subset.Red += count
		case "green":
			subset.Green += count
		}
	}
	return subset
}

func lineToGame(line string) Game {
	var id int
	var rest string

	_, _ = fmt.Sscanf(line, "Game %d", &id)
	rest = strings.TrimSpace(strings.Split(line, ":")[1])

	game := Game{
		ID:      id,
		Subsets: make([]Subset, 0),
	}
	setStrings := strings.Split(rest, ";")
	for _, set := range setStrings {
		game.Subsets = append(game.Subsets, setStrToSubset(set))
	}

	return game
}

func main() {
	lines := fileToLines("input.txt")
	//possibleIds := make([]int, 0, len(lines))
	powers := make([]int, 0, len(lines))
	for _, line := range lines {
		game := lineToGame(line)
		powers = append(powers, game.getPower())
	}
	sum := int(0)
	for _, n := range powers {
		sum += n
	}
	fmt.Println(sum)
}
