package main

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	set := make(map[rune]int, len(s))
	var lastAmount int
	var curAmount int

	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		j, found := set[char]

		if found {
			if curAmount > lastAmount {
				lastAmount = curAmount
			}
			curAmount = 0
			i = j + 1
			set = make(map[rune]int, len(s))
		}
		curAmount++
		char = rune(s[i])
		set[char] = i

	}
	if curAmount > lastAmount {
		lastAmount = curAmount
	}

	return lastAmount
}
