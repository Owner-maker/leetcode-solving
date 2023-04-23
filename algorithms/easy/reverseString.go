package main

func reverseString(s []byte) {
	length := len(s)

	if length == 1 {
		return
	}

	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}
