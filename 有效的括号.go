/*
思路概括：
1. 奇数长度直接返回 false（无法完全配对）。
2. 用栈存储左括号对应的期望右括号：遍历字符串，遇到左括号压入对应右括号，遇到右括号弹出栈顶比对。
3. 遍历结束后栈为空说明全部匹配。

复杂度：
- 时间：O(N)，遍历一次字符串。
- 空间：O(N)，最坏情况全为左括号时栈深度等于字符串长度。
*/

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
