//Write a Go function to find the first non-repeating character in a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func nonRepeat(s string) string {

	for _, val := range s {
		if strings.Count(s, string(val)) == 1 {
			return string(val)
		}
	}
	return "nil"
}
func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(nonRepeat(s))
}
