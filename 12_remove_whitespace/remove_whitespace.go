//Write a Go function to remove all whitespace from a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removeWhitecase(s string) string {

	removedSpace := ""
	for _, value := range s {
		if string(value) != " " {
			removedSpace += string(value)
		}
	}
	return removedSpace
}

func removeWhitecaseM2(s string) string {
	var result strings.Builder
	for _, char := range s {
		if char != ' ' {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(removeWhitecase(s))
	fmt.Println(removeWhitecaseM2(s))
}
