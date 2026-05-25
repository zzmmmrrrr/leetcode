package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	ans := 0
	curMin := prices[0]
	for _, p := range prices {
		curMin = min(p, curMin)
		ans = max(ans, p-curMin)
	}
	fmt.Println(ans)
}
