package main

import (
	"errors"
	"fmt"
	"log"
	"os"
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

func main() {
	fmt.Println("Starting...")
	content, err := os.ReadFile("input_file.txt")
	if err != nil {
		log.Fatal("Couldn't read file: input_file.txt. Please provide an input file.")
	}

	games := strings.Split(string(content), "\n")
	if len(games[len(games)-1]) == 0 {
		games = games[:len(games)-1]
	}
	total := addUpGameIds(games)

	if err != nil {
		log.Fatal("Couldn't calculate the totals because of: ", err)
	}

	fmt.Printf("Grand total: %d\n", total)
}

func addUpGameIds(inputLines []string) int {
	var totalGameIds int
	for _, inputLine := range inputLines {
		parsedGame, _ := parseGameData(inputLine)
		if isGamePossible(parsedGame) {
			totalGameIds += parsedGame.number
		}
	}
	return totalGameIds
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
