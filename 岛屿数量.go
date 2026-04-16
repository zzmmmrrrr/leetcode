package main

import (
	"fmt"
)

// 核心思想就是一旦遇到一个陆地地块，就不停dfs，想办法把连接的所有地块标记为2，这样就不会重复计算
func landsNum(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		for k := 0; k < 4; k++ {
			dfs(i+dx[k], j+dy[k])
		}

	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}
func main() {
	grid := [][]byte{
		{'1', '1', '1', '0'},
		{'1', '0', '0', '1'},
		{'0', '0', '0', '0'},
		{'1', '1', '1', '0'},
	}
	ans := landsNum(grid)
	fmt.Println(ans)
}
