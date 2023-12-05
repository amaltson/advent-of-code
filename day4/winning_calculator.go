package day4

import (
	"errors"
	"fmt"
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
	fmt.Println(inputLine[startNumIndex+1 : separatorIndex])
	fmt.Println(inputLine[separatorIndex+1:])
	whitespace := regexp.MustCompile(`\s+`)
	winningNumsStr := whitespace.Split(strings.TrimSpace(inputLine[startNumIndex+1:separatorIndex]), -1)
	cardNumsStr := whitespace.Split(strings.TrimSpace(inputLine[separatorIndex+1:]), -1)

	fmt.Printf("%v\n", winningNumsStr)
	fmt.Printf("%v\n", cardNumsStr)

	for _, winningNum := range winningNumsStr {
		fmt.Printf("Number: %s", winningNum)
		winning, err := strconv.Atoi(strings.TrimSpace(winningNum))
		if err != nil {
			return card{}, errors.New(`winning numbers are not numbers`)
		}
		fmt.Printf("Winnings: %d\n", winning)
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
