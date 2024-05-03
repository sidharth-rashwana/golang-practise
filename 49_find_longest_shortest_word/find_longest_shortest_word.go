// Write a Go function to find the shortest word in a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findShortest(s string) struct {
	name   string
	length int
} {
	words := strings.Fields(s)
	type word struct {
		name   string
		length int
	}

	var x word
	x.name = "none"
	x.length = 100

	for _, word := range words {
		if len(word) < x.length {
			x.name = word
			x.length = len(word)
		}
	}
	return x
}

func findLongest(s string) struct {
	name   string
	length int
} {
	words := strings.Fields(s)
	type word struct {
		name   string
		length int
	}

	var x word
	x.name = "none"
	x.length = -1000

	for _, word := range words {
		if len(word) > x.length {
			x.name = word
			x.length = len(word)
		}
	}
	return x
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(findShortest(s))
	fmt.Println(findLongest(s))
}
