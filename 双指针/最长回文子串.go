package main

import "fmt"

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}

func longestPalindrome(s string) (string, int) {
	if len(s) == 0 {
		return "", 0
	}

	bestStart := 0
	bestLen := 1

	for i := 0; i < len(s); i++ {
		start, end := expand(s, i, i)
		oddLen := end - start + 1
		if oddLen > bestLen {
			bestLen = oddLen
			bestStart = start
		}

		if i+1 < len(s) {
			start, end = expand(s, i, i+1)
			evenLen := end - start + 1
			if evenLen > bestLen {
				bestLen = evenLen
				bestStart = start
			}
		}

	}

	return s[bestStart : bestStart+bestLen], bestLen
}

func main() {

	s := "123329"
	substr, subLen := longestPalindrome(s)
	fmt.Println(substr, subLen)
}
