package main

import (
	"fmt"
)

// 快慢指针法
// 慢指针需要指向已经处理好的区域的边界
// 快指针每次移动一步，遇到需要处理的元素就进入处理逻辑
// 慢指针只有合法区域扩大时才移动
func solve(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
func main() {
	nums := []int{0, 0, 1, 0, 2, 3, 4, 0}
	fmt.Println(nums)
	solve(nums)
	fmt.Println(nums)
}
