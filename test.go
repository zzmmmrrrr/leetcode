package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	n := 5433
	nums := []int{1, 2, 3, 4, 5}

	// 预处理：排序是后续剪枝和贪心构造的前提
	sort.Ints(nums)
	maxDigit := nums[len(nums)-1]

	nStr := strconv.Itoa(n)
	lenN := len(nStr)
	nDigits := make([]int, lenN)
	for i, c := range nStr {
		nDigits[i] = int(c - '0')
	}

	bestSameLen := -1 // 实时维护同位数最优解，替代 candidates 数组
	found := false    // 配合提前终止机制

	var dfs func(curNum []int, idx int)
	dfs = func(curNum []int, idx int) {
		// 优化1: 如果已经找到合法解，由于我们是按从大到小/贪心顺序搜索
		// 第一个找到的必然是最大值，直接利用 found 标志位提前终止所有后续搜索
		if len(curNum) == lenN {
			temp := digitsToInt(curNum)
			if temp < n {
				bestSameLen = temp
			}
			return
		}

		targetDigit := nDigits[idx]

		// 优化2: 逆序遍历，优先尝试更大的数字，保证首次命中即为最优解
		for i := len(nums) - 1; i >= 0; i-- {
			x := nums[i]

			// 剪枝1：当前数字大于目标位，后续更小的数字才可能合法，跳过当前
			if x > targetDigit {
				continue
			}

			// 剪枝2：当前数字小于目标位，剩余位全填 maxDigit 即为该分支绝对最优解
			// 无需再向下递归枚举，直接计算并终止
			if x < targetDigit {
				best := make([]int, lenN)
				copy(best, curNum)
				best[idx] = x
				for j := idx + 1; j < lenN; j++ {
					best[j] = maxDigit
				}
				bestSameLen = digitsToInt(best)
				return
			}

			// x == targetDigit，保持紧约束，继续显式回溯向下搜索
			curNum = append(curNum, x)
			dfs(curNum, idx+1)
			curNum = curNum[:len(curNum)-1]

		}
	}

	dfs([]int{}, 0)

	// 优化4: Fallback 机制，同位数无解时 O(D) 构造少一位最大值
	if !found {
		if lenN <= 1 {
			fmt.Println("无解")
			return
		}
		fallback := make([]int, lenN-1)
		for i := range fallback {
			fallback[i] = maxDigit
		}
		bestSameLen = digitsToInt(fallback)
	}

	fmt.Println(bestSameLen)
}

func digitsToInt(digits []int) int {
	ans := 0
	for _, x := range digits {
		ans = ans*10 + x
	}
	return ans
}
