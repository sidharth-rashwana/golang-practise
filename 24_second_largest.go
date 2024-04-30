/*Write a program to find the second largest element in an array.*/
package main

import "fmt"

func secondLargest(s []int) int{
	for i:=0;i<len(s);i++{
		for j:=i+1;j<len(s);j++{
			if s[i] > s[j]{
			   s[i],s[j] = s[j],s[i]
		}
		}
	}
	return s[len(s)-2]
}

func main(){
	x :=[]int{5,1,3,2,4,8,7,6}
	secondLargest(x)
}