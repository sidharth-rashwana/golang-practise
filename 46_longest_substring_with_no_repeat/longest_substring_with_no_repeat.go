package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the string")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("you entered : ", str)
	substr := strings.Split(str, " ")
	longestSubstr := ""
	for _, val := range substr {
		if isUnique(val) {
			if len(val) > len(longestSubstr) {
				longestSubstr = val
			}
		}
	}
	fmt.Println("longest substring: ", longestSubstr)
}

// isUnique checks if a string has all unique characters
func isUnique(s string) bool {
	d := make(map[string]int)
	for _, val := range s {
		if _, exists := d[string(val)]; exists { // Check if the character already exists in the map
			return false
		}
		d[string(val)] = 1 // Mark this character as seen
	}
	return true
}
