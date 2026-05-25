package main

import (
	"fmt"
)

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	ans := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			ans = append(ans, left[i])
			i++
		} else {
			ans = append(ans, right[j])
			j++
		}
	}

	ans = append(ans, left[i:]...)
	ans = append(ans, right[j:]...)
	return ans
}

func main() {
	nums := []int{1, 9, 2, 1, 10, 23, -1}
	fmt.Println(mergeSort(nums))
}
