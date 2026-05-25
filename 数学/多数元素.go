package main

import "fmt"

func MooreVote(nums []int) int {
	voteCnt := 0
	candidateNum := 0

	for _, x := range nums {
		if voteCnt == 0 {
			candidateNum = x
		}
		if x == candidateNum {
			voteCnt++
		} else {
			voteCnt--
		}
	}
	return candidateNum

}

func hashMap(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
		if cnt[x] > len(nums)/2 {
			return x
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 3, 3, 3, 3, 3, 4, 4}
	a := MooreVote(nums)
	b := hashMap(nums)
	fmt.Print(a, b)
}
