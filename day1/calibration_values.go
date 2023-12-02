package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func addUpTheNumbers(allCalibrations []string) (int, error) {
	total := 0
	for _, line := range allCalibrations {
		if len(line) == 0 {
			continue
		}
		foundNumber, err := findFirstAndLastNumber(line)
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

	fmt.Printf("Grand total: %d", total)
}
