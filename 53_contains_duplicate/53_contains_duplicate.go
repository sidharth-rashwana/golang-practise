package main

import "fmt"

func containsDuplicate(nums []int) (bool, int) {
	x := make(map[int]int)
	for _, val := range nums {
		x[val] = x[val] + 1
	}
	for idx, val := range x {
		if val > 1 {
			return true, idx
		}
	}
	return false, 0
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
