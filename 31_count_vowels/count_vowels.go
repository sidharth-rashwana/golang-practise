/*Write a function to count the number of vowels in a given string.*/
package main

import "fmt"

func countVowels(s string) int {
	count := 0
	for _, val := range s {
		switch val {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count += 1
		}
	}
	return count
}
func main() {
	fmt.Println(countVowels("this is a sample text"))
}
