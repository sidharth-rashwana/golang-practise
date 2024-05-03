//Write a Go function to find the second most frequent character in a string.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func secondHighest(s string) {
	d := make(map[string]int)
	for _, val := range s {
		d[string(val)]++
	}
	// Create a slice of struct to store key-value pairs
	var keyValuePairs []struct {
		Key   string
		Value int
	}

	// Append key-value pairs from the map to the slice
	for k, v := range d {
		keyValuePairs = append(keyValuePairs, struct {
			Key   string
			Value int
		}{k, v})
	}

	// Sort the slice by values in descending order
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})

	// Extract keys from the sorted slice
	var sortedKeys []string
	for _, pair := range keyValuePairs {
		sortedKeys = append(sortedKeys, pair.Key)
	}
	fmt.Println(sortedKeys)
	fmt.Println(sortedKeys[1])

}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	secondHighest(s)
}
