// convert to lower and upper case

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertLower(s string) string {
	return strings.ToLower(s)
}

func convertUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	fmt.Println("Enter a string:")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(convertLower(s))
	fmt.Println(convertUpper(s))
}
