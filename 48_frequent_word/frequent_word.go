//Write a Go function to find the most frequent word in a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func frequentWord(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		m[word]++
	}
	return m
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(frequentWord(s))
}
