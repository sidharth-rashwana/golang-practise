package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	frequencies := make(map[string]int)
	words := strings.Fields(text) // Split the text into words

	// Iterate through each word and update its frequency
	for _, word := range words {
		frequencies[word]++
	}

	return frequencies

}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))

}
