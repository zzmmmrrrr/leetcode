/*
思路概括：
1. 随机选取 pivot 并与 left 交换，避免最坏情况退化。
2. partition 函数使用双指针将数组划分为小于 pivot 和大于 pivot 的两部分，返回 pivot 最终位置。
3. 递归对 pivot 左右两侧子数组排序。

复杂度：
- 时间：平均 O(N log N)，最坏 O(N²)（随机化可大幅降低最坏概率）。
- 空间：O(log N)，递归栈深度。
*/

package main

import (
	"fmt"
	"math/rand"
)

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	idx := left + rand.Intn(right-left+1)
	pivot := nums[idx]
	nums[idx], nums[left] = nums[left], nums[idx]
	i := left + 1
	j := right
	for {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

func main() {
	nums := []int{1, 9, 23, -2, -3, 1000}
	quickSort(nums, 0, len(nums)-1)
	for _, x := range nums {
		fmt.Println(x)
	}

}
