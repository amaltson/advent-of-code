package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	winningNums []int
	cardNums    []int
}

// Card   1:  9 32  7 82 10 36 31 12 85 95 |  7 69 23  9 32 22 47 10 95 14 24 71 57 12 31 59 36 68  2 82 38 80 85 21 92
func parseCard(inputLine string) (card, error) {
	var parsedWinningNums []int
	var parsedCardNums []int

	startNumIndex := strings.Index(inputLine, ": ")
	separatorIndex := strings.Index(inputLine, "| ")
	if startNumIndex == -1 || separatorIndex == -1 {
		return card{}, errors.New(`badly structured file, expecting format:
			'Card #: # # .. | # # ..'`)
	}
	whitespace := regexp.MustCompile(`\s+`)
	winningNumsStr := whitespace.Split(strings.TrimSpace(inputLine[startNumIndex+1:separatorIndex]), -1)
	cardNumsStr := whitespace.Split(strings.TrimSpace(inputLine[separatorIndex+1:]), -1)

	for _, winningNum := range winningNumsStr {
		winning, err := strconv.Atoi(strings.TrimSpace(winningNum))
		if err != nil {
			return card{}, errors.New(`winning numbers are not numbers`)
		}
		parsedWinningNums = append(parsedWinningNums, winning)
	}

	for _, cardNum := range cardNumsStr {
		cardInt, err := strconv.Atoi(strings.TrimSpace(cardNum))
		if err != nil {
			return card{}, errors.New(`winning numbers are not numbers`)
		}
		parsedCardNums = append(parsedCardNums, cardInt)
	}

	return card{winningNums: parsedWinningNums, cardNums: parsedCardNums}, nil
}

func calculatePoints(inputLine string) int {
	totalPoints := 0
	parsedCards, _ := parseCard(inputLine)
	for _, cardNum := range parsedCards.cardNums {
		if slices.Contains(parsedCards.winningNums, cardNum) {
			if totalPoints == 0 {
				totalPoints = 1
			} else {
				totalPoints = totalPoints * 2
			}
		}
	}
	return totalPoints
}

func calculateTotalPoints(cards []string) int {
	grandTotal := 0
	for _, card := range cards {
		grandTotal += calculatePoints(card)
	}
	return grandTotal
}

func main() {
	fmt.Println("Starting...")
	content, err := os.ReadFile("input_file.txt")
	if err != nil {
		log.Fatal("Couldn't read file: input_file.txt. Please provide an input file.")
	}

	cards := strings.Split(string(content), "\n")
	total := calculateTotalPoints(cards)

	if err != nil {
		log.Fatal("Couldn't calculate the totals because of: ", err)
	}

	fmt.Printf("Grand total: %d\n", total)
}
