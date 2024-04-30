/*4. Write a program to check if a given number is even or odd.*/
package main

import "fmt"

func checkEven(x int, y int) {
	if ((x + y) % 2) == 0 {
		fmt.Println("Even.")
	} else {
		fmt.Println("Odd.")
	}

}

func main() {

	var (
		x int
		y int
	)
	fmt.Println("Enter x and y: ")
	fmt.Scan(&x, &y)
	checkEven(x, y)
}
