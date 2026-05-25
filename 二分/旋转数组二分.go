package main

//数组是一个升序数组，在某个位置下标k发生了向左旋转
//也即从1，2，3，4，5变成了3，4，5，1，2
//数组允许有两部分有序，两部分内都是升序

//核心思路是根据端点值和nums[mid]的值比较来判断
//mid切分的两边哪一边是有序的
//此处的right和left的含义是target可能出现的区域的左右端点
//eg1: 4,5,1,2,3,nums[mid]=1,nums[left]=4,nums[right]=3,t=2
//eg2:3,4,5,6,1,2,nums[mid]=5,nums[left]=3,nums[right]=2,t=4
func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] <= nums[right] { //右侧升序有序
			if nums[right] >= target && nums[mid] <= target { //t在这个升序区间内
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { //左侧升序有序
			if nums[left] <= target && nums[mid] >= target { //t在这个降序区间内
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

	}
	return -1
}
