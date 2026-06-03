/*
思路概括：
1. 对数组排序，外层固定第一个数 nums[i]，内层用双指针 l、r 在 (i, n-1] 区间找两数之和为 -nums[i]。
2. 剪枝优化：
   - nums[i] > 0 直接 break（后面全正，无法凑零）。
   - nums[i] + nums[i+1] + nums[i+2] > 0 时 break（最小组合已超零）。
   - nums[i] + nums[n-2] + nums[n-1] < 0 时 continue（最大组合不足零，需增大第一个数）。
3. 找到三元组后跳过左右重复元素，避免结果重复。

复杂度：
- 时间：O(N²)，排序 O(N log N)，外层 O(N)，内层双指针 O(N)。
- 空间：O(log N) 排序栈空间（忽略结果数组）。
*/

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i+2 < n; i++ {
		//仅仅nums[i]就已经大于0
		//注意等于0时不应该剪枝，因为后续可以出现三个数字都为0
		if nums[i] > 0 {
			break
		}
		// //固定nums[i]时，当前最小组合已经大于0，代表后续不可能有合法解了
		// if nums[i]+nums[i+1]+nums[i+2] > 0 {
		// 	break
		// }
		// //固定nums[i]时，当前最大组合已经小于0，代表需要更换答案数组固定的第一个元素
		// if nums[i]+nums[n-2]+nums[n-1] < 0 {
		// 	continue
		// }

		//为了防止三元组的第一个元素出现重复,只考虑相同元素第一次出现的位置
		//注意这里不能通过考虑最后一次出现的位置来避免组合重复
		//原因是出现多个0的时候,000也是一组合法解
		//由于数组是升序排列的,只考虑最后一次出现的0,再从右侧区间找三元组后续更大的两个元素会导致漏掉000这组解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum == 0:
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				//跳过和左指针当前元素重复的元素
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
				//跳过和右指针当前元素重复的元素
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--

			case sum < 0:
				l++

			case sum > 0:
				r--
			}
		}

	}

	return ans

}
func main() {
	nums := []int{0, 0, 0, 2, -2, 1, 1, 3, -4}
	ans := threeSum(nums)

	for _, col := range ans {
		for _, x := range col {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}
