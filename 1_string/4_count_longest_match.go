package main

import (
	"fmt"
)

func findCommonChars(strings []string) string {
	charFreq := make(map[rune]int)

	for _, str := range strings {
		for _, char := range str {
			charFreq[char]++
		}
	}

	var commonChars string
	for char, freq := range charFreq {
		if freq == len(strings) {
			commonChars += string(char)
		}
	}

	return commonChars
}

func main() {
	strings := []string{"ABC", "ABCD", "ABCDEF", "ABD"}
	fmt.Println(findCommonChars(strings))
}
