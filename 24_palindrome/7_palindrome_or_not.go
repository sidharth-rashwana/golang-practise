/*Write a program to check if a given string is a palindrome.*/
package main

import (
	"fmt"
	"strings"
)

func checkPalindrome(s string) bool {
	s = strings.ToLower(s)
	total := len(s)
	i := 0
	j := total - 1
	for i < j {
		if s[i] != s[j] {
			fmt.Println("not same")
			return false
		}
		i += 1
		j -= 1
	}
	return true
}

func main() {
	var s string
	fmt.Println("Enter string:")
	fmt.Scan(&s)
	fmt.Println(checkPalindrome(s))
}
