//Write a Go function to count the number of characters in a string.

package main

import "fmt"

func countChar(s string) map[rune]int {
	m := make(map[rune]int)
	for _, value := range s {
		m[value]++
	}
	return m
}

func main() {
	var s string
	fmt.Scan(&s)
	m := countChar(s)
	for idx, val := range m {
		fmt.Printf("%c %d\n", idx, val)
	}
}