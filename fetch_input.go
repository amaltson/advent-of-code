package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	// Check if sufficient arguments are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run script.go <year> <day>")
		os.Exit(1)
	}

	year := os.Args[1]
	day := os.Args[2]

	// Construct the URL
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)

	// Format day with leading zero if necessary
	dayInt, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid day format. Day should be an integer.")
		os.Exit(1)
	}
	dayPadded := fmt.Sprintf("%02d", dayInt)

	// Create directory path
	dirPath := filepath.Join(year, "day"+dayPadded)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		panic(err)
	}

	// Full path for the input file
	filePath := filepath.Join(dirPath, "input_file.txt")

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Write the body to a file
	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File downloaded successfully: %s\n", filePath)
}
