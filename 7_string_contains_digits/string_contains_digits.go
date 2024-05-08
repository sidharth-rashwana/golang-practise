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
	// To know type of variable : fmt.Println(reflect.TypeOf(s))
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range digits {
		if !strings.Contains(s, string(val)) {
			return false
		}
	}
	return true
}

func onlyDigitsM2(s string) bool {
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
	fmt.Println(onlyDigitsM2(s))
}
