// Write a Go function to reverse a string.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseStringM1(s string) string {
	// traverse in reverse order
	length := len(s)
	rev := ""
	for j := length - 1; j >= 0; j-- {
		rev += string(s[j])
	}

	return string(rev)
}

func reverseStringM2(s string) string {
	// swap
	runes := []rune(s) //convert string into runes
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(reverseStringM1(s))
	fmt.Println(reverseStringM2(s))
}
