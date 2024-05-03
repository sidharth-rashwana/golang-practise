/*3. Implement a function to find the maximum of three numbers.*/

package main

import "fmt"

func find_greatest(x int, y int, z int) {
	greatest := x
	if y > greatest {
		greatest = y
	}
	if z > greatest {
		greatest = z
	}
	fmt.Printf("%v is the greatest. \n", greatest)

}
func main() {
	var (
		x int
		y int
		z int
	)
	fmt.Println("Enter x ,y and z:")
	fmt.Scan(&x, &y, &z)
	find_greatest(x, y, z)
}
