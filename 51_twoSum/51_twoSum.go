package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := []int{}
	for idx, val := range nums {
		t := target - val
		isTarget, i := findRemain(nums, t, idx)
		if isTarget {
			res = append(res, idx, i)
			return res
		}
	}
	return []int{}
}

func findRemain(nums []int, t, curValIdx int) (bool, int) {
	for idx, val := range nums {
		if val == t && idx != curValIdx {
			return true, idx
		}
	}
	return false, 0
}

func main() {
	x := []int{2, 7, 11, 15}
	fmt.Println(twoSum(x, 9))

	y := []int{3, 2, 4}
	fmt.Println(twoSum(y, 6))

	z := []int{3, 3}
	fmt.Println(twoSum(z, 6))
}
