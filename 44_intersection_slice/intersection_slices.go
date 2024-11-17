/*Create a function to find the intersection of two slices.*/

package main

import "fmt"

func slices(s1 []int , s2 []int) []int{
	m :=make(map[int]bool)
	common :=[]int{}

	for _, val := range s1 {
		m[val] = true
	}


	for _,val := range(s2){
		if m[val]{
			common = append(common,val)
		}
	}
	return common

}
func main(){
	s1 :=[]int{1,2,3,4}
	s2 :=[]int{2,3,5,8}
	fmt.Println("Intersecting values: ",slices(s1,s2))
}