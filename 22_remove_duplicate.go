/*Write a function to remove duplicate elements from a slice.*/

package main

import "fmt"

func removeDuplicate(s []int) []int{
	// create a map
	m := make(map[int]bool)
	// create a slice of type int
	unique := []int{}

	for _,value := range(s){
		if !m[value]{
			unique = append(unique,value)
			m[value] = true
		}
	}
	return unique
}

func main(){
	x := []int{1,2,3,1,55,3,22,312,33,2,65,69,80,80}
	fmt.Println("Unique values : ",removeDuplicate(x))
}