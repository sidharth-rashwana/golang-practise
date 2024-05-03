/*Create a function to generate the Fibonacci sequence up to n terms.*/
package main

import "fmt"

func fib(n int) int {
	if n == 1 {
		return n
	}
	return n * fib(n-1)
}
func main() {
	var x int
	fmt.Println("Enter a number : ")
	fmt.Scan(&x)
	fmt.Println(fib(x))
}
