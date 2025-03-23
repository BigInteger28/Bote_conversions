package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("engines13.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Skip the first line if it exists
	if scanner.Scan() {
		// First line is ignored
	}

	for scanner.Scan() {
		line := scanner.Text()

		// Find the second colon (:) in the line
		parts := strings.SplitN(line, ":", 3)
		if len(parts) < 3 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		// Extract the string after the second colon
		content := parts[2]

		// Ensure the content is long enough for the required indices
		if len(content) < 12 {
			fmt.Printf("Line too short for required indices: %s\n", line)
			continue
		}

		// Get the required characters: 3, 6, 9, and 12
		var result string
		if string(content[2]) == "5" {
			result = string(content[1]) + string(content[5]) + string(content[8]) + string(content[11])
		} else if string(content[5]) == "5" {
			result = string(content[2]) + string(content[4]) + string(content[8]) + string(content[11])
		} else if string(content[8]) == "5" {
			result = string(content[2]) + string(content[5]) + string(content[7]) + string(content[11])
		} else if string(content[11]) == "5" {
			result = string(content[2]) + string(content[5]) + string(content[8]) + string(content[10])
		} else {
			result = string(content[2]) + string(content[5]) + string(content[8]) + string(content[11])
		}

		// Append the result to the original line and print
		newLine := fmt.Sprintf("%s%s", line, result)
		fmt.Println(newLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Press a key to exit...")
	fmt.Scanln()
}
