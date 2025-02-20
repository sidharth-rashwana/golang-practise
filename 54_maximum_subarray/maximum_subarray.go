package main

import "fmt"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := 0

	for _, num := range nums {
		currentSum += num
		if currentSum > maxSum {
			maxSum = currentSum
		}
		if currentSum < 0 {
			currentSum = 0
		}
	}

	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxSubArray([]int{1, 2, -3, 4, 5}))
	fmt.Println(maxSubArray([]int{-1, -5, -3, 2, -7, 4, 5}))
}
