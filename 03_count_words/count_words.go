//Write a Go function to count the number of words in a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(s string) int {
	words := strings.Fields(s)
	return len(words) //total words
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(countWords(s))
}
