/*Implement a function to check if a number is a prime number.*/
package main

import "fmt"

func primeOrNote(n int) bool{
	if n ==0 || n == 1 || n == 2{
		return true
	}

	count :=0
	
	for i:=1;i<=n;i++{
		if n % i == 0 {
			count +=1
		}
		if count > 2{
			return false
		} 
	}
	return true
}

func main(){
	var n int
	fmt.Println("Enter a number : ")
	fmt.Scan(&n)
	checkPrime := primeOrNote(n)
	if checkPrime{
		fmt.Printf("The number %d is prime",n)
	}else{
		fmt.Printf("The number %d is not prime",n)
	}
}