//Write a Go function to find all permutations of a given string.

package main

import "fmt"

func permute(str string) []string {
	var result []string
	permuteHelper([]rune(str), 0, &result)
	return result
}

func permuteHelper(str []rune, index int, result *[]string) {
	if index == len(str)-1 {
		*result = append(*result, string(str))
		return
	}

	for i := index; i < len(str); i++ {
		str[index], str[i] = str[i], str[index] // Swap characters
		permuteHelper(str, index+1, result)
		str[index], str[i] = str[i], str[index] // Backtrack
	}
}

func main() {
	input := "abc"
	fmt.Println("Permutations of", input, "are:", permute(input))
}
