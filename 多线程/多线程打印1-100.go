/*
思路概括：
1. 三个协程通过三个 buffer=1 的 channel 进行令牌传递，轮流打印递增计数器 cnt（1-100）。
2. 协程 1 等待 chan1，打印后传递令牌给 chan2；协程 2 等待 chan2，传递给 chan3；协程 3 等待 chan3，传递回 chan1。
3. 当 cnt > 100 时，当前协程不打印，直接传递令牌给下一个协程并 return，形成连锁退出。
4. 主协程向 chan1 发送初始信号拨动第一个开关，通过 WaitGroup 等待所有协程完成。

复杂度：
- 时间：O(N)，每个数字打印一次。
- 空间：O(1)，仅使用固定数量的 channel。
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	k := 3
	chs := make([]chan struct{}, k)
	for i := 0; i < k; i++ {
		chs[i] = make(chan struct{}, 1)
	}
	var wg sync.WaitGroup
	wg.Add(k)
	curNum := 1
	for i := 0; i < k; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				<-chs[id]
				nextId := (id + 1) % k
				if curNum > 100 {

					chs[nextId] <- struct{}{}
					return
				}

				fmt.Printf("from chs[%d]:%d\n", id, curNum)
				curNum++
				chs[nextId] <- struct{}{}
			}
		}(i)
	}

	chs[0] <- struct{}{}
	wg.Wait()
	fmt.Println("finished")
}
