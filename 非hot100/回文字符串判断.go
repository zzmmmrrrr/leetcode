package main

import (
	"fmt"
)

func main() {
	s := "123321"
	t := "12322"

	fmt.Println(isPalindrome(s))
	fmt.Println(isPalindrome(t))
}

func isPalindrome(s string) bool {
	flag := true
	n := len(s)
	for i := 0; i < n/2; i++ {
		ch1, ch2 := s[i], s[n-1-i]
		if ch1 != ch2 {
			flag = false
			break
		}
	}
	return flag
}
