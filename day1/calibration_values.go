package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func main() {
	fmt.Println("Starting...")
	content, err := os.ReadFile("input_file.txt")
	if err != nil {
		log.Fatal("Couldn't read file: input_file.txt. Please provide an input file.")
	}

	total := 0
	calibrations := strings.Split(string(content), "\n")
	for _, line := range calibrations {
		foundNumber, _ := findFirstAndLastNumber(line)
		total += foundNumber
	}

	fmt.Printf("Grand total: %d", total)
}
