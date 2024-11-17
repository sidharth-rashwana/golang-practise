package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func longestShortest(s string) (int, int) {
	words := strings.Fields(s)
	var longest, shortest int
	longest = len(string(s[0]))
	shortest = len(string(s[0]))
	for _, val := range words {
		if longest < len(val) {
			longest = len(val)
		}
		if shortest > len(val) {
			shortest = len(val)
		}
	}
	return longest, shortest

}
func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(longestShortest(s))
}
