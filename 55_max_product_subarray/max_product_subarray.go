package main

import "fmt"

// Function to find the maximum product subarray
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd := nums[0]
	minProd := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			// Swap maxProd and minProd when a negative number is encountered
			maxProd, minProd = minProd, maxProd
		}

		// Update maxProd and minProd
		maxProd = max(nums[i], maxProd*nums[i])
		minProd = min(nums[i], minProd*nums[i])

		// Update the result with the maximum product so far
		result = max(result, maxProd)
	}

	return result
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to return the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums)) // Output: 6
}
