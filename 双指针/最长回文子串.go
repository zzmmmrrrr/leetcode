package main

import "fmt"

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - 1 - (left + 1) + 1 //退出for循环时right和left已经越界了，不满足s[left]==s[right]
}

func longest(s string) string {
	if len(s) == 0 {
		return ""
	}

	maxStart, maxLen := 0, 1
	for i := 0; i < len(s); i++ {
		odd := expand(s, i, i)
		even := expand(s, i, i+1)

		curLen := max(odd, even)
		if curLen > maxLen {
			maxLen = curLen
			maxStart = i - (curLen-1)/2
		}

	}
	return s[maxStart : maxStart+maxLen]

}

func main() {
	s := "12328"
	ans := longest(s)
	fmt.Println(ans)
}
