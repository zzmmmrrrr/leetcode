/*
思路概括：
1. 用双指针 i、j 分别遍历两个已排序数组。
2. 每次比较 nums1[i] 和 nums2[j]，将较小的元素追加到结果数组，并移动对应指针。
3. 一个数组遍历完毕后，将另一个数组剩余元素直接追加。

复杂度：
- 时间：O(M + N)，M、N 分别为两个数组的长度。
- 空间：O(M + N)，结果数组的长度。
*/

package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 3, 43, 4534, 22222}
	nums2 := []int{23, 66, 3333}
	ans := merge(nums1, nums2)
	for _, val := range ans {
		fmt.Println(val)
	}
}

func merge(nums1, nums2 []int) []int {
	i, j := 0, 0
	len1, len2 := len(nums1), len(nums2)
	ans := make([]int, 0, len1+len2)
	for i < len1 && j < len2 {
		if nums1[i] < nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}

	for i < len1 {
		ans = append(ans, nums1[i])
		i++
	}
	for j < len2 {
		ans = append(ans, nums2[j])
		j++
	}

	return ans

}
