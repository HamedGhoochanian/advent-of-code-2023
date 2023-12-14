package main

import (
	"fmt"
	"math"
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

type Card struct {
	Score   uint
	Matches uint
	Copies  uint
	Winning []int
	Mine    []int
}

func (c *Card) setScore() uint {
	power := -1

	for _, mine := range c.Mine {
		for _, winner := range c.Winning {
			if winner == mine {
				power++
			}
		}
	}
	var score uint
	if power >= 0 {
		score = uint(math.Pow(2.0, float64(power)))
	} else {
		score = 0
	}
	c.Matches = uint(power + 1)
	c.Score = score
	return score
}

func lineToCard(line string) Card {
	card := Card{
		Copies: uint(1),
	}
	a := strings.Split(line, ":")[1]
	split := strings.Split(a, "|")
	winning := strings.Split(split[0], " ")
	mine := strings.Split(split[1], " ")
	for _, val := range mine {
		if val != "" {
			parsed, _ := strconv.Atoi(val)
			card.Mine = append(card.Mine, parsed)
		}
	}
	for _, val := range winning {
		if val != "" {
			parsed, _ := strconv.Atoi(val)
			card.Winning = append(card.Winning, parsed)
		}
	}
	return card
}

func asd(cards []Card) {
	for i := range cards {
		for c := 0; c < int(cards[i].Copies); c++ {
			for j := 1; j <= int(cards[i].Matches); j++ {
				if i+j >= len(cards) {
					break
				}
				cards[i+j].Copies += 1
			}
		}
	}
}

func main() {
	lines := fileToLines("input.txt")
	cards := make([]Card, 0, len(lines))
	for _, line := range lines {
		card := lineToCard(line)
		card.setScore()
		cards = append(cards, card)
	}
	asd(cards)
	for _, card := range cards {
		fmt.Printf("%+v\n", card)
	}
	sum := 0
	for _, c := range cards {
		sum += int(c.Copies)
	}
	fmt.Println(sum)
}
