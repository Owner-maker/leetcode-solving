package main

import (
	"strings"
)

func reverseWords(s string) string {
	var builder strings.Builder

	runes := []rune(s)

	wordsEndings := make([]int, len(runes), len(runes))
	var j int

	for i, v := range runes {
		if v == ' ' {
			wordsEndings[j] = i - 1
			j++
		}
	}
	wordsEndings[j] = len(runes) - 1

	rightIndex := 0
	right := wordsEndings[rightIndex]

	for _, v := range runes {
		if v == ' ' {
			builder.WriteRune(' ')
			rightIndex++
			right = wordsEndings[rightIndex]
			continue
		}

		builder.WriteRune(runes[right])
		right--

	}

	return builder.String()
}
