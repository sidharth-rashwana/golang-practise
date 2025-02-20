package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a string : ")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	m := make(map[string]int)
	for _, val := range str {
		m[string(val)] += 1
	}
	fmt.Println(m)
}
