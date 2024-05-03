/*Implement a function to calculate the factorial of a number.*/
package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return n
	}
	return n * factorial(n-1)
}

func main() {
	var x int
	fmt.Println("Enter x:")
	fmt.Scan(&x)
	fmt.Println("Factorial of x :", x, "is", factorial(x))
}
