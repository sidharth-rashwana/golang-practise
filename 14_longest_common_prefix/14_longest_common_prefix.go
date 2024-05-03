//Write a Go function to find the longest common prefix string amongst an array of strings

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commonPrefix(s []string) string {
	// base condition
	if len(s) == 0 {
		return "none"
	}

	//find shortest
	shortest := s[0]
	for _, val := range s {
		if len(val) < len(shortest) {
			shortest = val
		}
	}

	final := ""
	for i := 0; i < len(shortest); i++ {
		ch := shortest[i] // get the current character
		for _, str := range s {
			if ch != str[i] {
				return final
			}
		}
		// if they are equal then add
		final += string(ch)
	}
	return final
}

func main() {
	var n int
	var r []string
	fmt.Println("Enter total sub strings:")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		s = strings.TrimSpace(s)
		r = append(r, s)
	}
	fmt.Println(commonPrefix(r))
}
