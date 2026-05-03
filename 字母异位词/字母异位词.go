/*
思路概括（两种方法）：
1. 排序法（groupAnagrams1）：将每个字符串排序后作为哈希表的键，排序后相同的字符串归入同一组。
2. 计数法（groupAnagrams2）：用长度为 26 的字节数组记录每个字母出现次数作为键分组（Go 中数组可比较，可直接作为 map key）。

复杂度：
- 排序法：时间 O(N × K log K)，空间 O(N × K)，N 为字符串个数，K 为平均长度。
- 计数法：时间 O(N × K)，空间 O(N × K)。
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"ate", "eat", "tea", "nike", "adi", "dia"}
	ans1 := groupAnagrams1(strs)
	ans2 := groupAnagrams2(strs)

	for _, x := range ans1 { //x是字符串切片
		for _, y := range x { //y是字符串
			fmt.Printf("%s ", y)
		}
		fmt.Println()
	}

	for _, x := range ans2 { //x是字符串切片
		for _, y := range x { //y是字符串
			fmt.Print(y, " ")
		}
		fmt.Println()
	}

}

//字母异位词指的是两个词字母组成相同，但是顺序不必一致
//比如tea和ate

// 1.方法1，使用排序后字符串为键+map的方法
// 原理是如果两个词互为异位词，那么两个词排序后一定相同
// 所以可以把排序后的词作为键，指向一个保存了互为异位词的字符串切片
func groupAnagrams1(strs []string) [][]string {
	mp := make(map[string][]string, len(strs))
	for _, oldStr := range strs {
		byteStr := []byte(oldStr)
		sort.Slice(byteStr, func(i, j int) bool {
			return byteStr[i] < byteStr[j]
		})
		sortedStr := string(byteStr)
		mp[sortedStr] = append(mp[sortedStr], oldStr)
	}
	ans := make([][]string, 0, len(mp))
	for _, val := range mp {
		ans = append(ans, val)
	}
	return ans
}

//2.方法2，使用字母计数数组作为键+map的方法
//原理是字母出现的范围固定，一定是a-z这26个字母
//可以用[26]byte来作为key，注意key必须是可比较类型
//在golang中数组可以比较，且数组的长度属于类型相关因素，比如[5]byte和[9]byte不是一个类型
//但是切片不能比较，因为切片是引用类型

func groupAnagrams2(strs []string) [][]string {
	mp := make(map[[26]byte][]string, len(strs))

	for _, s := range strs {
		cnt := [26]byte{}
		for i := 0; i < len(s); i++ {
			cnt[s[i]-'a']++
		}
		mp[cnt] = append(mp[cnt], s)
	}
	//注意这里要声明长度为0，容量为len(mp)
	//长度不为0会导致出现空切片
	ans := make([][]string, 0, len(mp))
	for _, val := range mp {
		ans = append(ans, val)
	}
	return ans

}
