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
