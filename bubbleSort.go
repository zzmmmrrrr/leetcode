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
