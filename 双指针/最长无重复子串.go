package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var s string

	fmt.Fscan(in, &s)
	cnt := make(map[byte]int)
	left := 0
	maxL := 1
	for right := 0; right < len(s); right++ {
		char := s[right]
		cnt[char]++
		for cnt[char] > 1 {
			cnt[s[left]]--
			left++
		}
		maxL = max(maxL, right-left+1)
	}
	fmt.Println(maxL)
}
