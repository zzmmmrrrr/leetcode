package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
)

/*
思路概括：
1. 将可用数字 nums 排序并去重，方便后续按从大到小尝试构造结果。
2. 把 n 拆分成数字数组 digitsN，按位进行 DFS 构造。
3. 对于当前位置：
   - 若选择的数字小于 n 当前位，则后续所有位直接填最大数字即可得到当前分支最优解（贪心剪枝）。
   - 若选择的数字等于 n 当前位，则继续递归处理下一位。
   - 若选择的数字大于 n 当前位，则该分支一定超过 n，直接跳过。
4. 如果最终无法构造出与 n 同位数且小于 n 的数字，则答案一定是长度比 n 少一位、且所有位都取最大数字的结果。

时间复杂度：
- 排序：O(m log m)，其中 m 为 nums 长度。
- DFS：
  由于一旦某位选到小于目标位的数字，就立即贪心填充后续位并返回，
  不会继续产生指数级搜索。
  每一位最多遍历 m 个数字，因此 DFS 总复杂度为 O(lenN × m)。
- 综合复杂度：O(m log m + lenN × m)。

空间复杂度：
- digitsN 数组占 O(lenN)。
- DFS 递归深度最多为 lenN，因此递归栈占 O(lenN)。
- 总空间复杂度：O(lenN)。
*/

func main() {
	n := 5433
	nums := []int{1, 2, 3, 4, 5}

	// 排序并去重，保证数字有序
	sort.Ints(nums)
	nums = slices.Compact(nums)

	// 将 n 转换为数字数组，方便逐位比较
	strN := strconv.Itoa(n)
	lenN := len(strN)
	digitsN := make([]int, lenN)
	for i, c := range strN {
		digitsN[i] = int(c - '0')
	}

	// 记录能够构造出的最大合法答案
	maxSameLenNum := -1

	/*
		curNum：当前已经构造出的前缀数字
		idx    ：当前处理到的位下标
	*/
	var dfs func(curNum []int, idx int)

	dfs = func(curNum []int, idx int) {
		// 已构造完成一个与 n 等长的数字
		if len(curNum) == lenN {
			temp := digitsToInt(curNum)

			// 必须严格小于 n
			if temp < n {
				maxSameLenNum = max(temp, maxSameLenNum)
			}
			return
		}

		targetDigit := digitsN[idx]

		// 从大到小枚举，优先尝试更大的数字
		for i := len(nums) - 1; i >= 0; i-- {

			// 当前位比目标位小
			if nums[i] < targetDigit {

				fmt.Printf("idx=%d, curNum before append=%v, cap=%d\n",
					idx, curNum, cap(curNum))

				curNum = append(curNum, nums[i])

				fmt.Printf("idx=%d, curNum after append=%v, cap=%d\n",
					idx, curNum, cap(curNum))

				/*
					关键剪枝：
					当前位已经严格小于 n 对应位，
					后面的位无需搜索，直接全部填最大数字即可得到
					当前分支的最优解。
				*/
				maxElement := nums[len(nums)-1]
				for j := idx + 1; j < lenN; j++ {
					curNum = append(curNum, maxElement)
				}

				maxSameLenNum = max(digitsToInt(curNum), maxSameLenNum)

				// 因为数字按降序遍历，第一个满足 < targetDigit 的方案
				// 已经是当前层能取得的最优结果，可以直接返回
				return
			}

			// 当前位比目标位大，一定超过 n，剪枝
			if nums[i] > targetDigit {
				continue
			}

			/*
				当前位与目标位相等，
				继续递归判断后续位是否能够构造出合法答案。
			*/
			curNum = append(curNum, nums[i])

			dfs(curNum, idx+1)

			// 回溯，恢复现场
			curNum = curNum[:len(curNum)-1]
		}
	}

	// 从最高位开始搜索
	dfs([]int{}, 0)

	if maxSameLenNum != -1 {
		fmt.Println(maxSameLenNum)
	} else {
		/*
			不存在与 n 同长度且小于 n 的合法数字。

			此时答案一定是长度少一位的最大数字，
			例如：
			n = 1000
			nums = [1,2,3]
			答案为 333
		*/
		maxShorterLenNum := make([]int, lenN-1)
		maxElement := nums[len(nums)-1]

		for i := range maxShorterLenNum {
			maxShorterLenNum[i] = maxElement
		}

		fmt.Println(digitsToInt(maxShorterLenNum))
	}
}

// 将数字数组转换成整数
func digitsToInt(digits []int) int {
	ans := 0
	for _, x := range digits {
		ans = ans*10 + x
	}
	return ans
}
