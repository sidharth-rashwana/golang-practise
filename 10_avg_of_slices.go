/*Write a function to calculate the average of a slice of floats.*/
package main

import "fmt"

func main() {

	num := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	total := 0.0
	for _, val := range num {
		total += val
	}
	fmt.Println(total / float64(len(num)))
}
