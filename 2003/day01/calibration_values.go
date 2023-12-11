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

func findNumberStart(calibration string) string {
	var foundNum string
	for i := 0; i < len(calibration); i++ {
		if _, err := strconv.Atoi(calibration[i : i+1]); err == nil {
			foundNum = calibration[i : i+1]
			return foundNum
		}
		for j := i; j <= len(calibration); j++ {
			if val, found := digitToIntMap[calibration[i:j]]; found {
				foundNum = val
				return foundNum
			}
		}
	}
	return foundNum
}

func findNumberEnd(calibration string) string {
	var foundNum string
	for i := len(calibration); i > 0; i-- {
		if _, err := strconv.Atoi(calibration[i-1 : i]); err == nil {
			foundNum = calibration[i-1 : i]
			return foundNum
		}
		for j := i; j > 0; j-- {
			if val, found := digitToIntMap[calibration[j:i]]; found {
				foundNum = val
				return foundNum
			}
		}
	}
	return foundNum
}

func findFirstAndLastNumber(calibration string) (int, error) {
	startNumber := findNumberStart(calibration)
	endNumber := findNumberEnd(calibration)
	foundNumbers, err := strconv.Atoi(startNumber + endNumber)
	return foundNumbers, err
}

func addUpTheNumbers(allCalibrations []string) (int, error) {
	total := 0
	for _, line := range allCalibrations {
		if len(line) == 0 {
			continue
		}
		startNumber := findNumberStart(line)
		endNumber := findNumberEnd(line)
		foundNumbers, _ := strconv.Atoi(startNumber + endNumber)
		total += foundNumbers
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
