package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	runes := []rune(s)
	l := len(runes)

	for i := 0; i < l/2; i++ {
		if runes[i] != runes[l-i-1] {
			return false
		}
	}

	return true
}
