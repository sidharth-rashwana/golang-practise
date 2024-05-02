package main

import (
	"fmt"
	"sort"
)

func sortByValue(m map[string]int) []string {
	// Create a slice of struct to store key-value pairs
	var keyValuePairs []struct {
		Key   string
		Value int
	}

	// Append key-value pairs from the map to the slice
	for k, v := range m {
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

	return sortedKeys
}

func main() {
	// Example map
	m := map[string]int{
		"apple":  3,
		"banana": 2,
		"cherry": 5,
		"date":   1,
	}

	// Sort the map by value
	sortedKeys := sortByValue(m)

	// Print the sorted keys
	fmt.Println("Sorted keys by value:")
	for _, key := range sortedKeys {
		fmt.Println(key)
	}
}
