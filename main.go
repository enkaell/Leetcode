package main

import (
	"fmt"
)

func isValid(s string) bool {
	right := len(s) - 1
	left := 0
	for left < len(s)-1 {
		if int(s[right])-int(s[left]) != 1 {
			right--
			fmt.Println(left, right)
		} else {
			fmt.Println(left, right)
			left++
			right = len(s) - 1
			fmt.Printf("%s %s \n", s[left], s[right])
		}
	}
	return false
}
func main() {
	isValid("({})")
}
