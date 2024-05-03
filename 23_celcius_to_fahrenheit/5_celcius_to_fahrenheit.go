/*5.Create a program to convert temperature from Celsius to Fahrenheit.*/

package main

import "fmt"

func cToF(c float64) float64 {
	F := (9.0/5.0)*c + 32
	fmt.Println(F)
	return F
}

func main() {
	var (
		c float64
	)

	fmt.Println("Enter temperature in c:")
	fmt.Scanf("%f", &c)
	fmt.Println(cToF(c))
}
