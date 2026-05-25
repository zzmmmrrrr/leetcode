/*
思路概括：
1. DP 方法：预处理 lMax[i]（i 及其左侧最大高度）和 rMax[i]（i 及其右侧最大高度），每个位置 i 的接水量为 min(lMax[i], rMax[i]) - height[i]。
2. 双指针方法：用 i、j 从两端向中间靠拢，维护当前左右最大高度 lMax、rMax。哪边最大高度更小，则该侧当前位置的接水量已确定（对侧只会更高），累加后移动该侧指针。

复杂度：
- 时间：O(N)，均只需一次遍历。
- 空间：DP 方法 O(N)（两个辅助数组），双指针方法 O(1)。
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{6,2,3,4,5,6}
	ans := trap(nums)
	fmt.Println("这是dp方法", ans)
	fmt.Println("这是双指针方法，状态压缩", twoPointers(nums))
}
func trap(height []int) int {
	n := len(height)
	ans := 0
	lMax, rMax := make([]int, n), make([]int, n)
	//lMax[i],rMax[i]表示角标i位置元素自身以及左侧和右侧最高点
	lMax[0] = height[0]
	rMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		lMax[i] = max(lMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(rMax[i+1], height[i])
	}
	//每个位置最多能接雨水的高度是左右板子中的短板减去地基高度
	for i, h := range height {
		ans += min(lMax[i], rMax[i]) - h
	}
	return ans

}


//双指针法
//把dp数组进行状态压缩，O(n)的空间复杂度压缩到O(1)
//哪边的 max 更小，哪边的结果就已经确定了，因为对面只会更大，不影响短板的判断。
func twoPointers(height []int) int{
	ans:=0
	n:=len(height)
	i,j:=0,n-1
	lMax,rMax:=0,0
	for i<=j{
		lMax=max(lMax,height[i])
		rMax=max(rMax,height[j])

		if lMax<rMax{
			ans+=lMax-height[i]
			i++
		} else{
			ans+=rMax-height[j]
			j--
		}
	}
	return ans
}