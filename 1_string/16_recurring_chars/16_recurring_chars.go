//Write a Go function to find the first recurring character in a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findRecurringChars(s string) rune {

	m := make(map[string]int)

	for _, val := range s {
		m[string(val)]++
		if m[string(val)] > 1 {
			return val
		}
	}
	return 0
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	output := string(findRecurringChars(s))
	fmt.Println(output)
	if output == string(0) {
		fmt.Println("No recurring character")
	} else {
		fmt.Println("Recurring character: ", string(output))
	}
}
