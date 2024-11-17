package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) []rune {
	x := []rune(s) // convert string into slice
	length := len(x)
	for i := 0; i < length/2; i++ {
		x[i], x[length-i-1] = x[length-i-1], x[i]
	}
	return x
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(string(reverseString(s)))
}
