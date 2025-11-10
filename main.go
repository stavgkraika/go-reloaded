package main

import (
	"fmt"
	"os"
)

func runAllTests() {
	fmt.Println("Running tests...")
	// Add test implementation here if needed
	fmt.Println("Tests completed.")
}

func main() {
	// Check for test mode
	if len(os.Args) == 2 && os.Args[1] == "test" {
		runAllTests()
		return
	}

	// Validate command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		fmt.Println("       go run . test")
		return
	}

	// Read input file
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Process text and write output
	processor := NewProcessor(string(input))
	output := processor.Process()

	err = os.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
}