/*
思路概括：
1. 方法一（两个辅助数组）：用 left[i] 表示 nums[i] 左侧所有元素的乘积，right[i] 表示右侧所有元素乘积，ans[i] = left[i] * right[i]。
2. 方法二（滚动变量优化空间）：先用 ans 存储 left 值，再从右向左遍历用变量 right 滚动累乘右侧乘积，直接更新 ans[i]。

复杂度：
- 时间：O(N)，两次遍历。
- 空间：方法一 O(N)（两个辅助数组），方法二 O(1)（仅使用输出数组，不计入额外空间）。
*/

package main

import (
	"fmt"
)

// 1.方法1，用两个额外的数组来存储左右侧局部乘积
func solve1(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	//left[i],right[i]表示nums[i]左侧和右侧所有元素乘积，不包含nums[i]
	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = 1, 1

	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		ans[i] = left[i] * right[i]
	}
	return ans

}

//2.方法2，最优解
//先用ans[i]存储left[i],再用right作为滚动数组，right来表示right[i]，随着i的变化滚动更新

func solve2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	right := 1 //现在的right相当于right[i-1]
	for i := n - 1; i >= 0; i-- {
		ans[i] *= right
		right *= nums[i]
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	ans1 := solve1(nums)
	ans2 := solve2(nums)
	for _, x := range ans1 {
		fmt.Println(x)
	}
	for _, x := range ans2 {
		fmt.Println(x)
	}
}
