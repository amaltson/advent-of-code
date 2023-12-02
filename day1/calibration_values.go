package main

import (
	"fmt"
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
}
