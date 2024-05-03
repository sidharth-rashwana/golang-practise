//Write a Go function to find all duplicate characters in a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findDuplicates(s string) string {
	m := make(map[rune]int)
	for _, val := range s {
		m[val]++
	}

	r := ""
	for key, value := range m {
		if value > 1 {
			r += string(key)
		}
	}
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	s = strings.TrimSpace(s)
	duplicates := findDuplicates(s)
	fmt.Println("Duplicates:", duplicates)
}
