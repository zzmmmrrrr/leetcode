/*
思路概括：
1. 对 digits 排序去重，利用贪心思想从高位到低位匹配 n 的每一位。
2. 优先尝试构造长度等于 n 的位数的数字：
   - 若当前位能选相等的数字，则递归向下。
   - 若向下失败，则回溯并在当前位选一个更小的数字，随后后面所有位直接填最大数字。
3. 若同位数构造失败（如 digits 最小数都比 n 的第一位大），则返回长度为 len(n)-1 且全由 digits 最大值组成的数。

复杂度：
- 时间：O(K log K + L * K)，K 为 digits 长度，L 为 n 的位数。
- 空间：O(L)，用于存储位数数组及递归栈。
*/

package main

import (
	"fmt"
	"sort"
)

// maxNumLessThanN 寻找由 digits 组成的、小于 n 的最大整数
func maxNumLessThanN(digits []int, n int) int {
	// 1. 预处理：排序、去重，并将 n 拆解为每一位
	sort.Ints(digits)
	digits = unique(digits)
	nToDigits := intToDigits(n)
	l := len(nToDigits)
	ans := make([]int, l)

	// 2. 尝试构造一个与 n 位数相同的数字
	x := solve(digits, nToDigits, ans, 0)
	if x != -1 {
		return x
	}
	/*	如果要求严格小于，可以在外部函数里面修改
		保持solve函数的逻辑统一性，也即不需要单独判断是否在处理最后一位且元素相等
		if x!=-1 && x!=n{
			return x
		}
	*/

	// 3. 失败处理：如果无法构造同位数数字（例如 n=2111, digits={2,9}）
	// 则退而求其次，构造一个比 n 少一位的最大数，每一位都填 digits 里的最大值
	maxDigit := digits[len(digits)-1]
	y := 0
	for i := 0; i < l-1; i++ {
		y = y*10 + maxDigit
	}

	return y
}

// solve 通过 DFS 尝试构造与 n 同位数的数
func solve(digits, nToDigits, ans []int, idx int) int {
	lenN := len(nToDigits)
	// 递归终止条件：已经填满了所有位数，说明构造出了一个等于 n 的数（在题目要求“小于”的情况下，此处逻辑会被调用层微调）
	if idx == lenN {
		return digitsToInt(ans)
	}

	// 贪心：从大到小尝试 digits 里的数字
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i]

		// 如果当前数字大于 n 对应位的数字，不能选，继续看更小的
		if d > nToDigits[idx] {
			continue
		}

		// 如果当前数字小于 n 对应位的数字
		if d < nToDigits[idx] {
			// 贪心确定：当前位一旦变小，后面所有位必须填最大值才能保证整体最大
			ans[idx] = d
			maxDigit := digits[len(digits)-1]
			for j := idx + 1; j < lenN; j++ {
				ans[j] = maxDigit
			}
			return digitsToInt(ans)
		}

		// 如果d等于 n 对应位的数字
		//需要根据后续位置递归的结果进行判断
		ans[idx] = d
		postPart := solve(digits, nToDigits, ans, idx+1)

		// 如果后面位数的递归找到了合法解，直接返回
		if postPart != -1 {
			return postPart
		}

		// 如果 postPart == -1，说明当前位选 d (等于 n[idx]) 导致后面没法构造了
		// 循环会继续，尝试比 d 更小的数字（即 i--）
	}

	//只有Digits中每个数字都大于n的当前位数字时，才会退出for循环来到此处
	return -1 // 当前位尝试了所有数字都无法构造，返回失败
}

// intToDigits 将整数转换为切片，例如 5432 -> [5, 4, 3, 2]
func intToDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := []int{}
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}
	// 因为是取模获取，需要反转得到正常顺序
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// digitsToInt 将切片转换回整数，例如 [5, 4, 3, 2] -> 5432
func digitsToInt(digits []int) int {
	ans := 0
	for _, val := range digits {
		ans = ans*10 + val
	}
	return ans
}

// unique 对有序切片进行原地去重
func unique(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	ans := []int{nums[0]}
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			ans = append(ans, nums[i])
		}
	}
	return ans
}

func main() {
	cases := []struct {
		digits []int
		n      int
		expect int
	}{
		{[]int{1, 3, 5, 9}, 5432, 5399},
		{[]int{1, 2, 9}, 2151, 2129},
		{[]int{2, 9}, 2111, 999},
		{[]int{8, 9}, 88, 88},
	}

	for _, c := range cases {
		got := maxNumLessThanN(c.digits, c.n)
		fmt.Printf("digits=%-10v n=%-6d 期望=%-6d 实际=%-6d %s\n",
			c.digits, c.n, c.expect, got,
			map[bool]string{true: "✓", false: "✗"}[got == c.expect])
	}
}
