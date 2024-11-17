package main

import (
	"fmt"
	"sort"
)

func sortSlice(x []int) []int {
	sort.SliceStable(x, func(i, j int) bool {
		return x[i] < x[j]
	})
	return x
}
func main() {
	x := []int{1000, 2321, 322, 11233, 23, 231, 2321, 31, 23, 123, 12, 3123, 123123123123}
	fmt.Println(sortSlice(x))
}
