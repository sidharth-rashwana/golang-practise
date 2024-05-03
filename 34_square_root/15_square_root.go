/*Write a program to find the square root of a given number.*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	sqrt := math.Sqrt(num)
	fmt.Println(sqrt)
}
