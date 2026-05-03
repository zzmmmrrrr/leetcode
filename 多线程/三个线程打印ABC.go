/*
思路概括：
1. 三个协程分别打印 "A"、"B"、"C"，通过三个缓冲为 1 的 channel 形成令牌环：chA → chB → chC → chA。
2. 每个协程循环 cnt 轮，等待自己的 channel 信号，打印字符后向下一个 channel 发送信号。
3. 主协程向 chA 发送初始信号激活，通过 WaitGroup 等待所有协程完成。

复杂度：
- 时间：O(cnt)，每个协程打印 cnt 次。
- 空间：O(1)，仅三个固定 channel。
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	const cnt = 10

	//创建三个channel帮助传递信号
	//预留缓冲区，避免最后一轮chC向chA传递令牌时，chA已关闭
	//导致chC阻塞卡死
	chA := make(chan struct{}, 1)
	chB := make(chan struct{}, 1)
	chC := make(chan struct{}, 1)

	//使用wg启动，这里的wg.Add必须在协程启动之前完成
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < cnt; i++ {
			<-chA
			fmt.Print("A")
			chB <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < cnt; i++ {
			<-chB
			fmt.Print("B")
			chC <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < cnt; i++ {
			<-chC
			fmt.Print("C")
			chA <- struct{}{}
		}
	}()

	chA <- struct{}{}

	wg.Wait()
	fmt.Println()
}
