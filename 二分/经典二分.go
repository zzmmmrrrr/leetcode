package main

import "fmt"

//返回插入位置的idx
func binarySearch(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	ans := binarySearch(nums, target)
	fmt.Println(ans)
}
