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
	for idx, _ := range shortest {
		ch := shortest[idx] // get the current character
		for _, str := range s {
			// if not matching then only return
			if ch != str[idx] {
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
