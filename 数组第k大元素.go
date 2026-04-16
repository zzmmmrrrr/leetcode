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
