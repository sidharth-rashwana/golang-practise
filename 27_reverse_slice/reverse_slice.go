/*Implement a function to reverse a slice of integers.*/

package main

import "fmt"

func reverseSlice(s []int) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	reverseSlice(x)
	fmt.Println(x)
}
