/*
思路概括：
1. 外层循环 i 每轮为第 i 个位置选出未排序部分的最小值。
2. 内层循环从 i+1 开始，若发现更小值则交换，并标记 isChanged = true。
3. 若某轮未发生任何交换，说明数组已有序，提前终止。

复杂度：
- 时间：最坏 O(N²)，最好 O(N)（已有序时提前退出）。
- 空间：O(1)，原地排序。
*/

package main

import (
	"fmt"
)

func bubbleSort(nums []int) {

	n := len(nums)
	for i := 0; i < n; i++ {
		isChanged := false
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}
func main() {
	nums := []int{1, 4, -1, -2, 47, -100, 23}
	bubbleSort(nums)
	for _, x := range nums {
		fmt.Println(x)
	}
}
