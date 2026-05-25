/*
思路概括：
1. 遍历数组，用哈希表记录每个数的索引。
2. 对当前数 num，检查 target - num 是否已在哈希表中，若在则直接返回两个索引。
3. 若未出现则将 num 及其索引存入哈希表，继续遍历。

复杂度：
- 时间：O(N)，仅需一次遍历。
- 空间：O(N)，哈希表存储已遍历的元素。
*/

package main

import "fmt"

//如果x+y==target,返回x，y的index
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int) //key是num，val是idx，检查num是否存在
	for i, num := range nums {
		diff := target - num
		if _, isAppeared := mp[diff]; isAppeared {
			return []int{mp[diff], i} //注意这里返回的是mp[diff]
		}

		mp[num] = i
	}
	return []int{}
}
func main() {
	test := []int{1, -1, 2, 3, 4, 45, 66, -8, 22}
	target := 0
	ans := twoSum(test, target)

	for _, x := range ans {
		fmt.Println(x)
	}
}
