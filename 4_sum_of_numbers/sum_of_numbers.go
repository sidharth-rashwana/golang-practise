/*Create a program to calculate the sum of two numbers.*/

package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}
func main() {
	/*
		var (
			x int
			y int
		)
	*/
	var x, y int
	fmt.Println("Enter x and y :")
	fmt.Scan(&x, &y)
	c := sum(x, y)
	fmt.Println("Total: ", c)
}
