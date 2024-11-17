package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findRecurringChars(s string) string {

	m := make(map[rune]int)
	for _, val := range s {
		m[val]++
		if m[val] > 1 {
			return string(val)
		}
	}
	return "All characters occured only once."
}

func main() {
	fmt.Println("Enter a string : ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(findRecurringChars(s))
}
