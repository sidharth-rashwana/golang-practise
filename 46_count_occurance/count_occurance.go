//Write a Go function to count the number of occurrences of each character in a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countOccurance(s string) map[string]int {
	m := make(map[string]int)
	for _, val := range s {
		m[string(val)]++
	}
	return m

}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(countOccurance(s))
}
