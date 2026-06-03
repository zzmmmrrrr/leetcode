package main

import "fmt"

func main() {
	nums := []int{}
	ans := longestContinous(nums)
	fmt.Print(ans)
}

func longestContinous(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := 1
	n := len(nums)
	set := make(map[int]bool, n)
	for _, x := range nums {
		set[x] = true
	}

	for _, x := range nums {
		if set[x-1] { //x-1也在nums中,x不应该作为最长连续序列的起点
			continue
		}

		steps := 1
		for set[x+steps] {
			steps++
		}
		ans = max(ans, steps)
		//注意进入到这一步是x+steps已经不在set里面了,此时的steps已经越界了,超出连续序列了
		//因此序列是从x 到x+steps-1,长度是steps-1+1=steps

	}
	return ans

}
