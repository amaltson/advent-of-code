package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	number int
	cubes  map[string][]int
}

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func parseGameData(inputLine string) (game, error) {
	re := regexp.MustCompile(`Game (\d+): (.*)`)
	matches := re.FindStringSubmatch(inputLine)
	cardNumber, err := strconv.Atoi(matches[1])
	if err != nil {
		return game{}, errors.New("couldn't extract card number")
	}
	resultingCubes := map[string][]int{}
	cubeSetsLine := matches[2]
	cubeSets := strings.Split(cubeSetsLine, "; ")
	for _, cubeSet := range cubeSets {
		colourBalls := strings.Split(cubeSet, ", ")
		for _, ball := range colourBalls {
			numberAndBall := strings.Split(ball, " ")
			colour := numberAndBall[1]
			num, _ := strconv.Atoi(numberAndBall[0])
			resultingCubes[colour] = append(resultingCubes[colour], num)
		}
	}
	return game{number: cardNumber, cubes: resultingCubes}, nil
}

func isGamePossible(checkingGame game) bool {
	for colour, maxOfColour := range maxCubes {
		for _, cubeNum := range checkingGame.cubes[colour] {
			if cubeNum > maxOfColour {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("Hello, world!")
}
