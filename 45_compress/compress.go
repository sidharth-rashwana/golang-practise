//cccbbdaaacca --> c3b2d1a3c2a1

package main

import (
	"fmt"
	"strconv"
)

func compressString(s string) string {
	count := 1
	evaluated := ""
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
		} else {
			evaluated += string(s[i]) + strconv.Itoa(count)
			count = 1
		}
	}
	// append last char
	evaluated += string(s[len(s)-1]) + strconv.Itoa(count)
	return evaluated
}

func main() {
	fmt.Println(compressString("cccbbdaaacca"))
}
