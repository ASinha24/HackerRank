package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	rev := []rune(s)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return string(rev)
}

func main() {
	fmt.Println("reverse of a string")
	fmt.Println(strings.Repeat("_", 45))
	str := "still love you"
	fmt.Println(reverseString(str))
}
