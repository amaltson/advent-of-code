package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingGameData(t *testing.T) {
	assert := assert.New(t)

	example := `Game 1: 1 red, 3 blue, 11 green; 1 blue, 5 red; 3 blue, 5 green, 13 red; 6 red, 1 blue, 4 green; 16 red, 12 green`
	expected := map[string][]int{
		"red":   []int{1, 5, 13, 6, 16},
		"blue":  []int{3, 1, 3, 1},
		"green": []int{11, 5, 4, 12},
	}
	game, err := parseGameData(example)
	assert.Equal(1, game.number)
	assert.Equal(expected, game.cubes)
	assert.Nil(err)
}

func TestAdditionalParsingGameData(t *testing.T) {
	assert := assert.New(t)

	example := `Game 14: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red`
	expected := map[string][]int{
		"red":   []int{3, 6, 14},
		"blue":  []int{6, 15},
		"green": []int{1, 3, 3},
	}
	game, err := parseGameData(example)
	assert.Equal(14, game.number)
	assert.Equal(expected, game.cubes)
	assert.Nil(err)
}

func TestGameIsPossible(t *testing.T) {
	assert := assert.New(t)

	possibleGame := game{
		number: 4,
		cubes: map[string][]int{
			"red":   []int{1, 5, 4},
			"blue":  []int{3, 1, 3, 1, 10},
			"green": []int{11, 5, 4, 12},
		},
	}

	assert.True(isGamePossible(possibleGame))
}

func TestGameIsNotPossible(t *testing.T) {
	assert := assert.New(t)

	possibleGame := game{
		number: 4,
		cubes: map[string][]int{
			"red":   []int{1, 5, 13, 6, 16},
			"blue":  []int{3, 1, 3, 1},
			"green": []int{11, 5, 4, 12},
		},
	}

	assert.False(isGamePossible(possibleGame))
}
