package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findFirstNumber(line string) (int, error) {
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
	file, err := os.Open("input_file.txt")
	if err != nil {
		log.Fatal("Couldn't open file: input_file.txt. Please provide an input file.")
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		foundNumber, _ := findFirstNumber(line)
		total += foundNumber
	}

	fmt.Printf("Grand total: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal("Ran into an error with the bufio Scanner")
	}
}
