/*Implement a program to find the minimum element in an array.*/

package main

import "fmt"

func main() {
	x := []int{8, 9, 10, 7, 6, 2, 3, 1, 7, 55, 44, 4}
	l := x[0]
	for _, val := range x {
		if val < l {
			l = val
		}
	}
	fmt.Println("least number:", l)
}
