package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] != s[l-i-1] {
			fmt.Printf("%c %c", s[i], s[l-i-1])
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(isPalindrome(s))
}
