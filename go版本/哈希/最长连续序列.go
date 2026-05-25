package main

import (
	"fmt"
)

func main() {
	var ans int
	nums := []int{1, 2, 2, 3, 5, 6, 7, 8, 9, 10, 8}
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
	for num := range set { //num是一个数字
		if _, isAppeared := set[num-1]; isAppeared {
			continue //num-1在set中，那他就不是最优的起点
		}
		//当前数字可以作为起点
		//步长从1开始
		// 最终break时代表nums[num+steps-1]在集合中
		//但是nums[num+steps]不在，所以长度是steps-1+1=steps
		//这里的长度相当于是数组元素格式
		steps := 1
		for {
			if _, isAppeared := set[num+steps]; isAppeared {
				steps++
			} else {
				break
			}
		}
		ans = max(steps, ans)

	}
	fmt.Println(ans)
}
