// Write a Go function to check if a string contains only lowercase letters.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isLower(s string) bool {

	for _, val := range s {
		if !unicode.IsLower(val) {
			return false
		}
	}
	return true

}

func isUpper(s string) bool {

	for _, val := range s {
		if !unicode.IsUpper(val) {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println("Enter string:")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(isLower(s))
	fmt.Println(isUpper(s))
}
