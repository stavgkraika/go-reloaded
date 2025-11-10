package main

import (
	"fmt"
	"os"
)

func runTest(testNum int, input, expected string) {
	processor := NewProcessor(input)
	result := processor.Process()
	
	status := "PASS"
	if result != expected {
		status = "FAIL"
	}
	
	fmt.Printf("Test %d: %s\n", testNum, status)
	if status == "FAIL" {
		fmt.Printf("  Input:    %q\n", input)
		fmt.Printf("  Expected: %q\n", expected)
		fmt.Printf("  Got:      %q\n", result)
	}
}

func runAllGoldenTests() {
	fmt.Println("Running all golden tests...")
	
	// Test 1: Basic hexadecimal conversion
	runTest(1, "1E (hex)", "30")
	
	// Test 2: Binary conversion
	runTest(2, "101 (bin)", "5")
	
	// Test 3: Uppercase single word
	runTest(3, "Go (up)", "GO")
	
	// Test 4: Lowercase range
	runTest(4, "HELLO WORLD (low, 2)", "hello world")
	
	// Test 5: Capitalization range
	runTest(5, "brooklyn bridge (cap, 2)", "Brooklyn Bridge")
	
	// Test 6: Article correction before vowel
	runTest(6, "A orange tree.", "An orange tree.")
	
	// Test 7: Quotes cleanup
	runTest(7, "He said: ' this is fine ' .", "He said: 'this is fine'.")
	
	// Test 8: Punctuation spacing
	runTest(8, "Wait ,what ?!", "Wait, what?!")
	
	// Test 9: Ellipsis and punctuation handling
	runTest(9, "I was thinking ... You were right", "I was thinking... You were right")
	
	// Test 10: Out-of-range uppercase modifier
	runTest(10, "10 files (up, 6)", "10 FILES")
	
	// Test 11: Combined hex and uppercase
	runTest(11, "1f (hex) files (up)", "31 FILES")
	
	// Test 12: Quotes and command together
	runTest(12, "He shouted ' stop (up) '", "He shouted 'STOP'")
	
	// Test 13: Capitalization and punctuation
	runTest(13, "new york (cap) !", "New York!")
	
	// Test 14: Mixed uppercase and lowercase
	runTest(14, "one two three (up, 2) four (low)", "one TWO THREE four")
	
	// Test 15: Binary with punctuation
	runTest(15, "101 (bin)!", "5!")
	
	// Test 16: Punctuation spacing cleanup
	runTest(16, "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!")
	
	// Test 17: Empty quotes handling
	runTest(17, "He said: ''", "He said: ''")
	
	// Test 18: Article correction before 'h'
	runTest(18, "A honest mistake.", "An honest mistake.")
	
	// Test 19: Mixed modifiers with punctuation
	runTest(19, "wow (up), amazing (low, 1)!", "WOW, amazing!")
	
	// Test 20: Long sentence with multiple rules
	runTest(20, "I added 1E (hex) files ,and 10 (bin) were removed (low, 2) .", "I added 30 files, and 2 were removed.")
	
	// Test 21: Article correction before consonant
	runTest(21, "An lemon tree.", "A lemon tree.")
	
	// Test 22: Long sentence with multiple rules and article correction
	runTest(22, "harold wilson (cap, 2) : ' I am a (cap) optimist ,but a (up) optimist who carries , ' a raincoat .", "Harold Wilson: 'I Am An optimist, but An optimist who carries,' a raincoat.")
	
	fmt.Println("All tests completed.")
}

func init() {
	if len(os.Args) > 1 && os.Args[1] == "golden" {
		runAllGoldenTests()
		os.Exit(0)
	}
}