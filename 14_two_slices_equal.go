/* Implement a function to check if two slices are equal.*/

package main

import (
	"fmt"
)

func sliceEqual(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for idx, _ := range s1 {
		if s1[idx] != s2[idx] {
			return false
		}
	}
	return true

}

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{1, 2, 3, 4}
	fmt.Println(sliceEqual(s1, s2))
}
