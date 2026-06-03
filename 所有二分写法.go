package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 4, 4, 4, 5, 6}
	fmt.Printf("数组: %v (长度=%d)\n", nums, len(nums))
	fmt.Println("==============================================================")
	fmt.Printf("%-8s | %-12s | %-12s | %-12s\n", "Target", "LowerBound", "UpperBound", "LB(target+1)")
	fmt.Println("--------------------------------------------------------------")

	for target := 0; target <= 7; target++ {
		lb := lowerBound(nums, target)
		ub := upperBound(nums, target)
		lbNext := lowerBound(nums, target+1)

		// 格式化输出，对齐列便于肉眼对比
		fmt.Printf("%-8d | %-12d | %-12d | %-12d\n", target, lb, ub, lbNext)
	}
	fmt.Println("==============================================================")
}

// =============================================================================
// 【1】lowerBound：找第一个 >= target 的元素索引
// =============================================================================
// 作用：返回数组中第一个「大于等于」target 的位置。
// 场景：
//   - 精确查找的前置校验（找到后检查该位置的值是否真的 == target）
//   - 统计 target 出现次数的左边界
//   - 等价于 C++ std::lower_bound / Python bisect.bisect_left
// 注意：若所有元素都 < target，返回 -1（你的原始设计），实际工程中常返回 len(nums)
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := -1 // 未找到时返回 -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// mid 太小，答案一定在右半区
			left = mid + 1
		} else {
			// nums[mid] >= target：mid 是一个候选答案
			// ✅ 关键：不停止！记录 ans 后继续向左收缩，寻找「更早」的合法位置
			ans = mid
			right = mid - 1
		}
	}
	return ans
}

// =============================================================================
// 【2】upperBound：找第一个 > target 的元素索引
// =============================================================================
// 作用：返回数组中第一个「严格大于」target 的位置。
// 场景：
//   - 统计 target 出现次数的右边界（个数 = upperBound - lowerBound）
//   - 查找「不超过 target 的最大元素」的下一个位置
//   - 等价于 C++ std::upper_bound / Python bisect.bisect_right
// 注意：若所有元素都 <= target，返回 -1
func upperBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := -1 // 未找到时返回 -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			// mid 不够大，答案一定在右半区
			left = mid + 1
		} else {
			// nums[mid] > target：mid 是一个候选答案
			// ✅ 关键：不停止！记录 ans 后继续向左收缩，寻找「更早」的合法位置
			ans = mid
			right = mid - 1
		}
	}
	return ans
}

// =============================================================================
// 【3】biSearch（标准精确查找）：找到任意一个 == target 的位置
// =============================================================================
// 作用：确认 target 是否存在，存在则返回其某个索引。
// 场景：
//   - 只需判断「有没有」，不关心是第几个
//   - LeetCode 704 等基础二分题
// ⚠️ 陷阱：遇到重复元素时，返回的索引不确定（取决于 mid 落在哪）
//   例如 [4,4,4,4] 查 4，可能返回 1、2 或 3，绝不保证是最左或最右
func biSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid // ⚠️ 立即返回，不再探索左右边界
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // 未找到返回 -1（原代码返回 left，改为 -1 语义更明确）
}
