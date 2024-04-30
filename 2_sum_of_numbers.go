/*2. Create a program to calculate the sum of two numbers.*/

package main

import "fmt"

func main() {
	var (
		x int
		y int
	)
	fmt.Println("Enter x and y :")
	fmt.Scan(&x, &y)
	c := x + y
	fmt.Println("Total: ", c)
}
