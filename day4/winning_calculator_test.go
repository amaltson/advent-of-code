package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberExtractor(t *testing.T) {
	assert := assert.New(t)

	inputLine := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	card := parseCard(inputLine)
	assert.Equal([]int{41, 48, 83, 86, 17}, card.winningNums)
	assert.Equal([]int{83, 86, 6, 31, 17, 9, 48, 53}, card.cardNums)
}
