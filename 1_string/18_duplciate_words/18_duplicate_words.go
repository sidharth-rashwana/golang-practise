//Write a Go function to find all duplicate words in a string.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findDuplicateWords(s string) []string {
	duplicate := []string{}
	words := strings.Fields(s)
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		if count > 1 {
			duplicate = append(duplicate, word)
		}
	}

	return duplicate
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	duplicateWords := findDuplicateWords(s)
	fmt.Println("Duplicate words:", duplicateWords)
}
