/*【函数实现说明】
1. maxNumLessThanN(digits []int, n int) int
   - 功能：寻找由给定数字集合 digits 组成的、严格小于 n 的最大整数。
   - 参数：digits (可选数字切片), n (限制上限的目标整数)。

2. dfs(digits, nToDigits, ans []int, idx int) bool
   - 功能：通过 DFS 贪心回溯，尝试构造一个与 n 位数相同且严格小于 n 的最大数。
   - 参数：digits (去重排序后的可选数字), nToDigits (n 的逐位切片), ans (当前构造中的数字切片), idx (当前处理的位索引)。

3. intToDigits(n int) []int / digitsToInt(digits []int) []int
   - 功能：实现整数与逐位切片（如 5432 <-> [5, 4, 3, 2]）之间的相互转换。

4. unique(nums []int) []int
   - 功能：对已经有序的切片进行原地去重。

【复杂度分析】
- 时间复杂度：O(K log K + L * K)
  其中 K 为 digits 的长度，L 为 n 的十进制位数。
  - 排序去重耗时 O(K log K)。
  - DFS 过程中，每一位最多遍历整个 digits 数组（长为 K），总共有 L 位，因此构造同位数耗时 O(L * K)。
- 空间复杂度：O(L)
  - 存储 nToDigits 和 ans 切片各占用 O(L) 空间。
  - DFS 递归栈的最大深度为 L，占用 O(L) 空间。
*/

package main

import (
	"fmt"
	"slices"
	"sort"
)

func maxNumLessThanN(digits []int, n int) int {
	//1.预处理：排序去重digits，拆解n为位数组
	sort.Ints(digits)
	digits = unique(digits)
	digitsN := intToDigits(n)
	lenN := len(digitsN)
	digitsAns := make([]int, lenN)

	if dfs(digits, digitsN, digitsAns, 0) {
		return digitsToInt(digitsAns)
	}

	maxDigit := digits[len(digits)-1]
	ans := 0
	for i := 0; i < lenN-1; i++ {
		ans = ans*10 + maxDigit
	}
	return ans
}

func dfs(digits []int, digitsN []int, digitsAns []int, idx int) bool {
	lenN := len(digitsN)
	if idx == lenN {
		return true
	}
	//如果能走到idx==lenN，代表前面的位没有触发d<digitsN[idx]，因此此时一定是ans=n
	//如果题目允许返回小于等于n的最大值，此处是true
	//如果不允许就返回false

	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i]

		if d < digitsN[idx] {
			digitsAns[idx] = d
			maxDigit := digits[len(digits)-1]
			for j := idx + 1; j < lenN; j++ {
				digitsAns[j] = maxDigit
			}
			return true
		} else if d > digitsN[idx] {
			continue
		} else {
			digitsAns[idx] = d
			if dfs(digits, digitsN, digitsAns, idx+1) {
				return true
			}
		}

	}
	return false
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
	slices.Reverse(res)
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
	n := 5432
	digits := []int{5, 4, 3, 2}
	ans := maxNumLessThanN(digits, n)
	fmt.Println(ans)
}
