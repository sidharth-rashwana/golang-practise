// 3. Write a Go function to reverse a string.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) string {
	runes := []rune(s)//convert string into runes
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(reverseString(s))
}
