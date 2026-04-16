package main

import (
	"fmt"
)

func main() {
	test := "()[]{}"
	fmt.Println(isValid(test))
}
func isValid(s string) bool {
	if len(s)&1 == 1 {
		return false
	} //奇数长度
	stk := []byte{}
	mp := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for i := 0; i < len(s); i++ { //遇到左符号把右符号压进去，
		if val, ok := mp[s[i]]; ok {
			stk = append(stk, val)
		} else { //遇到右符号出栈比对
			if len(stk) == 0 || stk[len(stk)-1] != s[i] {
				return false
			}
			stk = stk[:len(stk)-1]
		}
		
	}
	return len(stk) == 0
}
