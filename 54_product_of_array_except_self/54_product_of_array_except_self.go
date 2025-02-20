package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	leftProd := 1
	for i := 0; i < n; i++ {
		result[i] = leftProd
		leftProd *= nums[i]

		// 1st iteration (index = 0)
		// result[0] = 1 (no elements to the left)
		// leftProd = 1 * 1 = 1

		// 2nd iteration (index = 1)
		// result[1] = 1 (leftProd from previous iteration)
		// leftProd = 1 * 2 = 2

		// 3rd iteration (index = 2)
		// result[2] = 2 (leftProd from previous iteration)
		// leftProd = 2 * 3 = 6

		// 4th iteration (index = 3)
		// result[3] = 6 (leftProd from previous iteration)
		// leftProd = 6 * 4 = 24
	}

	fmt.Println(result) // [1 1 2 6]

	rightProd := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProd
		rightProd *= nums[i]

		// 1st iteration (index = 3)
		// result[3] = 6 * 1 = 6
		// rightProd = 1 * 4 = 4

		// 2nd iteration (index = 2)
		// result[2] = 2 * 4 = 8
		// rightProd = 4 * 3 = 12

		// 3rd iteration (index = 1)
		// result[1] = 1 * 12 = 12
		// rightProd = 12 * 2 = 24

		// 4th iteration (index = 0)
		// result[0] = 1 * 24 = 24
		// rightProd = 24 * 1 = 24

	}

	return result // [24, 12, 8, 6]
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result)
}
