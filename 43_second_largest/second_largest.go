/*Write a program to find the second largest element in an array.*/
package main

import (
	"fmt"
	"sort"
)

func secondLargest(s []int) int {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	fmt.Println(s)
	return s[len(s)-2]
}

func main() {
	x := []int{5, 1, 3, 2, 4, 8, 7, 6}
	fmt.Println(secondLargest(x))
}
