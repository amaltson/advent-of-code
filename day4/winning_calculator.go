package day4

import (
	"fmt"
	"regexp"
)

type card struct {
	winningNums []int
	cardNums    []int
}

func parseCard(inputLine string) card {
	parsedWinningNums := []int{}
	parsedCardNums := []int{}

	numRe := regexp.MustCompile(`Card\s+\d:([\s\d]+) |([\s\d]+)`)
	matches := numRe.FindAllStringSubmatch(inputLine, -1)
	fmt.Printf("All matches: %v", matches)
	fmt.Printf("All matches: %s", matches[0])
	fmt.Printf("Match 1: %s", matches[1])
	fmt.Printf("Match 2: %s", matches[2])

	return card{winningNums: parsedWinningNums, cardNums: parsedCardNums}
}
