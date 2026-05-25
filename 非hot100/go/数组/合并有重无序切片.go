/*
思路概括：
1. 用 map 模拟集合对两个切片去重：将两个切片的元素全部插入 map，键为数值，值为空结构体。
2. 将 map 中的键收集到结果切片中，调用 sort.Ints 排序返回。

复杂度：
- 时间：O(M + N + K log K)，M、N 为两切片长度，K 为去重后元素个数。最坏 O((M+N) log (M+N))。
- 空间：O(K)，来自集合和最终数组。最坏 O(M + N)。
*/
package main

import (
	"fmt"
	"sort"
)

func solve(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	ans := make([]int, 0, m+n)
	set := make(map[int]struct{}, m+n)
	for _, num := range nums1 {
		set[num] = struct{}{}
	}
	for _, num := range nums2 {
		set[num] = struct{}{}
	}
	for key := range set {
		ans = append(ans, key)
	}

	sort.Ints(ans)
	return ans
}
func main() {
	a := []int{1, 2, -2, 1, 2, -3, -4}
	b := []int{9, 23, 11, 1, 1, 1, 1}
	ans := solve(a, b)
	for _, val := range ans {
		fmt.Print(val, " ")
	}
}
