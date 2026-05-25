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
	for i := 0; i < len(s); i++ { //遇到左符号把左符号压栈
		if _, ok := mp[s[i]]; ok {
			stk = append(stk, s[i])
		} else { //遇到右符号出栈比对
			if len(stk) == 0 {
				return false
			}
			top := stk[len(stk)-1]
			if mp[top] != s[i] {
				return false
			}

			stk = stk[:len(stk)-1]
		}

	}
	return len(stk) == 0
}
