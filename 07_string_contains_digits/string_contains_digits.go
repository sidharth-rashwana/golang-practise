//Write a Go function to check if a string contains only digits.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func onlyDigits(s string) bool {
	for _, val := range s {
		if !unicode.IsDigit(val) {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(onlyDigits(s))
}
