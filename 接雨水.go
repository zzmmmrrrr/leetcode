package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 2, 0, 3, 2, 5}
	ans := trap(nums)
	fmt.Println("这是dp方法", ans)
	fmt.Println("这是双指针方法，状态压缩", trap1(nums))
}
func trap(height []int) int {
	n := len(height)
	ans := 0
	lMax, rMax := make([]int, n), make([]int, n)
	//某个角标位置元素自身以及左侧和右侧最高点
	lMax[0] = height[0]
	rMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		lMax[i] = max(lMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(rMax[i+1], height[i])
	}
	for i, h := range height {
		ans += min(lMax[i], rMax[i]) - h
	}
	return ans

}

func trap1(height []int) int {
	n := len(height)
	ans := 0
	i, j := 0, n-1
	lMax, rMax := 0, 0
	for i < j {
		//lMax = 左指针 i 及其左侧的最高值
		//rMax = 右指针 j 及其右侧的最高值
		lMax = max(lMax, height[i])
		rMax = max(rMax, height[j])

		//这里判断的核心逻辑是，当lMax<rMax时，他一定小于右侧全局最大，所以可以直接那lmax当边界
		if lMax < rMax {
			ans += lMax - height[i]
			i++
		} else {
			ans += rMax - height[j]
			j--
		}

	}
	return ans
}
