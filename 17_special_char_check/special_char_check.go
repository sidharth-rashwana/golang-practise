// Write a Go function to check if a string contains any special characters.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkSpecial(s string) bool {
	specialChars := "!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	for _, char := range specialChars {
		if strings.Contains(s, string(char)) {
			return true
		}
	}
	return false
}
func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(checkSpecial(s))
}
