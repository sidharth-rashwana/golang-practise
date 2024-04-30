/*Implement a program to find the largest element in a 2D array.*/

package main

import "fmt"

func largestElement(arr2d [][]int) int{
	largest := arr2d[0][0]
	for i:=0;i<len(arr2d);i++{
		for j:=0;j<len(arr2d);j++{
			if arr2d[i][j] > largest{
				largest = arr2d[i][j]
			}
		}
	}
	return largest
}
func main(){
	var arr2d = [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	l := largestElement(arr2d)
	fmt.Println(l)
}