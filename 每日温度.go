/*
题目：给定一个温度数组，每个值代表当天的温度，要求返回一个数组，
每个值代表所在位置要等待几天才能遇到温度更高的日期
*/
package main

import (
	"fmt"
)

func main() {
	cases := [][]int{
		{73, 74, 75, 71, 69, 72, 76, 73},
		{30, 40, 50, 60},
		{30, 60, 90},
	}
	for _, c := range cases {
		fmt.Println(dailyTemperatures(c))
	}
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	//维护一个单调栈，不断把当前没有遇到温度更高的日期的idx加入栈中
	//这样单调栈就是一个温度单调递减的日期idx组成的栈
	//只有某天温度大于栈顶温度时，才弹出栈顶
	stack := make([]int, 0, n)

	for i, curTemp := range temperatures {
		for len(stack) > 0 {
			topIdx := stack[len(stack)-1]
			if curTemp > temperatures[topIdx] {
				ans[topIdx] = i - topIdx
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}

	return ans
}
