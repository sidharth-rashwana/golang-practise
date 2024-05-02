//Write a Go function to replace all occurrences of a substring within a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replaceSubstringM1(s string, sub string, replaceWith string) string {
	replacedString := ""

	subLen := len(sub)
	strLen := len(s)

	for idx := 0; idx < strLen; {
		// Check if the remaining length is less than the length of the substring
		if idx+subLen > strLen {
			// If yes, add the remaining characters to the result string and break
			replacedString += s[idx:]
			break
		}

		// Check if the current substring matches the target sub
		if s[idx:idx+subLen] == sub {
			// If yes, add replaceWith to the result string
			replacedString += replaceWith
			// Move the index past the length of the substring
			idx += subLen
		} else {
			// If no, add the current character to the result string
			replacedString += string(s[idx])
			// Move to the next character
			idx++
		}
	}

	return replacedString
}

func replaceSubstringM2(s string, sub string, replaceWith string) string {
	return strings.ReplaceAll(s, sub, replaceWith)
}

func main() {
	fmt.Println("Enter string:")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println("Enter sub string:")
	sub, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	sub = strings.TrimSpace(sub)
	fmt.Println("Enter replace with string:")
	replace, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	replace = strings.TrimSpace(replace)
	fmt.Println(replaceSubstringM1(s, sub, replace))
	fmt.Println(replaceSubstringM2(s, sub, replace))

}
