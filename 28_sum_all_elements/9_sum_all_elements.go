/*Create a program to find the sum of all elements in an array.*/

package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	total := 0
	for _, val := range x {
		total += val
	}
	fmt.Println(total)
}
