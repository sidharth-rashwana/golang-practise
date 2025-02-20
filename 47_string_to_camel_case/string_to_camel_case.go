// Write a Go function to convert a string to a camel case.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("enter the string")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// Trim spaces and split the string into words based on spaces and non-alphabetic characters
	words := strings.Fields(str)
	for idx, word := range words {
		if idx > 0 {
			// Capitalize the first letter of each word (except for the first word)
			words[idx] = strings.Title(strings.ToLower(word))
		} else {
			// Keep the first word in lowercase
			words[idx] = strings.ToLower(word)
		}
	}
	fmt.Println(strings.Join(words, ""))
}

// Input: "hello world this is go"

// Output: "helloWorldThisIsGo"
