package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter a string: ")
	var str string
	// Read the full input (including spaces) using Scanln, which captures everything until Enter is pressed.
	fmt.Scanln(&str)
	// Use strings.TrimSpace to remove leading and trailing spaces
	str = strings.TrimSpace(str)
	fmt.Println("Resulting string:", str, "Length:", len(str))
}
