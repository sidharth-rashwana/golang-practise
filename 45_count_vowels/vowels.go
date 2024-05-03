//  Write a Go function to count the number of vowels in a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countVowels(s string) map[string]int {
	v := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	m := make(map[string]int)
	for _, val := range s {
		for _, value := range v {
			if string(value) == string(val) {
				m[string(val)]++
				break
			}
		}
	}
	return m
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(countVowels(s))
}
