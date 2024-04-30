//6. Write a Go function to check if two strings are anagrams of each other.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func stringSort(s string) []rune {
	runes := []rune(s) // convert to rune

	// sort the rune
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// return the rune
	return runes
}

func isAnagram(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	string1 := stringSort(s1)
	string2 := stringSort(s2)

	for idx, _ := range string1 {
		if string1[idx] != string2[idx] {
			return false
		}
	}
	return true

}

func main() {

	fmt.Println("Enter string s1: ")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	s1 = strings.TrimSpace(s1)

	fmt.Println("Enter string s2: ")
	reader2 := bufio.NewReader(os.Stdin)
	s2, _ := reader2.ReadString('\n')
	s2 = strings.TrimSpace(s2)

	fmt.Println(isAnagram(s1, s2))
}
