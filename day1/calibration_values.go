package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var digitToIntMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func findFirstAndLastNumber(line string) (int, error) {
	// find first number from front
	var firstNum string
	var secondNum string
	for i := 0; i < len(line); i++ {
		if _, err := strconv.Atoi(line[i : i+1]); err == nil {
			firstNum = line[i : i+1]
			break
		}
	}
	for i := len(line); i > 0; i-- {
		if _, err := strconv.Atoi(line[i-1 : i]); err == nil {
			secondNum = line[i-1 : i]
			break
		}
	}
	return strconv.Atoi(firstNum + secondNum)
}

func convertSpelledOutNumbers(calibration string) string {
	for i := 0; i < len(calibration); i++ {
		for j := i; j <= len(calibration); j++ {
			if val, found := digitToIntMap[calibration[i:j]]; found {
				firstPart := calibration[:i]
				remainingPart := calibration[j:]
				calibration = firstPart + val + remainingPart
			}
		}
	}
	return calibration
}

func addUpTheNumbers(allCalibrations []string) (int, error) {
	total := 0
	for _, line := range allCalibrations {
		if len(line) == 0 {
			continue
		}
		normalizedLine := convertSpelledOutNumbers(line)
		foundNumber, err := findFirstAndLastNumber(normalizedLine)
		if err != nil {
			return -1, err
		}
		total += foundNumber
	}
	return total, nil
}

func main() {
	fmt.Println("Starting...")
	content, err := os.ReadFile("input_file.txt")
	if err != nil {
		log.Fatal("Couldn't read file: input_file.txt. Please provide an input file.")
	}

	calibrations := strings.Split(string(content), "\n")
	total, err := addUpTheNumbers(calibrations)

	if err != nil {
		log.Fatal("Couldn't calculate the totals because of: ", err)
	}

	fmt.Printf("Grand total: %d\n", total)
}
