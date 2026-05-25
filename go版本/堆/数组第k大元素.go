/*
思路概括：
1. 利用快速选择的 partition 分治思想，将第 K 大转化为第 (n-k) 小的索引 targetIdx。
2. 每次 partition 随机选取 pivot，将数组划分为小于和大于 pivot 的两部分。
3. 若 pivot 位置等于 targetIdx 则返回该值；若 pivot 偏右则向左区间搜索，否则向右区间搜索。

复杂度：
- 时间：平均 O(N)，最坏 O(N²)（随机化大幅降低最坏概率）。
- 空间：O(1)，迭代版本，原地操作。
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{8, 4, 2, 5, 1, 3, 2, 5}
	ans := findKth(nums, 1)
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(ans)
	for _, x := range nums {
		fmt.Println(x)
	}
}
func findKth(nums []int, k int) int {
	n := len(nums)
	targetIdx := n - k
	//第k大的元素是第n-k+1小的元素，也即落在n-k索引上
	//比如n=5，k=2,第2大也是第4小
	left, right := 0, n-1
	for {
		pivot := partition(nums, left, right)
		if pivot == targetIdx {
			return nums[pivot]
		}
		if pivot > targetIdx {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

}

func partition(nums []int, left int, right int) int {
	i := left + rand.Intn(right-left+1)
	//Intn会生成[0,right-left]的值，加上left就是[left,right]
	pivot := nums[i]
	nums[i], nums[left] = nums[left], nums[i]
	i = left + 1
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
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}
