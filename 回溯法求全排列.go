package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permutations(nums))
}

func permutations(nums []int) [][]int {
	n := len(nums)
	room := 1
	for i := 2; i <= n; i++ {
		room *= i
	}
	res := make([][]int, 0, room)
	path := make([]int, 0, n)
	used := make([]bool, n)

	var dfs func()
	dfs = func() {
		if len(path) == n {
			tmp := slices.Clone(path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true
			dfs()

			path = path[:len(path)-1]
			used[i] = false
		}
	}

	dfs()
	return res
}
