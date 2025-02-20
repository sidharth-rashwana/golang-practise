package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter a string :")
	var str string
	fmt.Scanf("%s", &str) // single input
	if len(str) == 0 {
		fmt.Println("empty string")
		return
	}
	fmt.Println("upper case : ", strings.ToUpper(str))
	fmt.Println("lower case : ", strings.ToLower(str))
}
