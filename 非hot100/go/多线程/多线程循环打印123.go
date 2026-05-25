/*
思路概括：
1. buffer 模式：创建 n 个缓冲为 1 的 channel，每个协程等待自己的 channel，打印 id+1 后取模传递给下一个协程。主协程通过 chans[0] 发送初始信号。
2. quit 模式：使用无缓冲 channel，最后一个协程在最后一次打印后不再传递 token，实现优雅退出。

复杂度：
- 时间：O(R × N)，R 为打印轮数，N 为协程数。
- 空间：O(N)，N 个 channel。
*/

package main

import (
        "fmt"
        "sync"
)

func buffer(n int) {
        // 1. 初始化通道切片
        // 注意：是一维切片 []chan struct{}
        chans := make([]chan struct{}, n)
        for i := 0; i < n; i++ {
                // 每个元素都需要 make，否则是 nil，读写都会死锁
                chans[i] = make(chan struct{},1)
        }

        var wg sync.WaitGroup
        wg.Add(n)

        maxRounds := 10 // 每个协程打印 10 次

        // 2. 循环启动 n 个协程
        for i := 0; i < n; i++ {
                go func(id int) {
                        defer wg.Done()
                        for j := 0; j < maxRounds; j++ {
                                // 等待属于自己的信号
                                <-chans[id]

                                // 打印逻辑：这里打印 id+1，即 1, 2, 3...
                                fmt.Print(id + 1)

                                // 计算下一个协程的索引：通过取模实现环形传递
                                next := (id + 1) % n
                                chans[next] <- struct{}{}
                        }
                }(i)
        }

        // 3. 激活第一个：拨动第一个开关
        chans[0] <- struct{}{}

        wg.Wait()
        fmt.Println("\n任务完成")
}

func quit(n int) {
        // 1. 初始化无缓冲通道
        chans := make([]chan struct{}, n)
        for i := 0; i < n; i++ {
                chans[i] = make(chan struct{}) // 无缓冲
        }

        var wg sync.WaitGroup
        wg.Add(n)

        maxRounds := 10 // 每人打印 10 次

        // 2. 启动协程
        for i := 0; i < n; i++ {
                go func(id int) {
                        defer wg.Done()
                        for j := 0; j < maxRounds; j++ {
                                // 等待信号
                                <-chans[id]

                                fmt.Print(id + 1)

                                // 【核心退出逻辑】
                                // 判断：如果这是最后一个协程（id == n-1）
                                // 且这是最后一轮打印（j == maxRounds-1）
                                // 那么就不再向第一个通道（chans[0]）发送信号，直接退出。
                                if id == n-1 && j == maxRounds-1 {
                                        return
                                }

                                // 否则，正常把棒传给下一个人
                                next := (id + 1) % n
                                chans[next] <- struct{}{}
                        }
                }(i)
        }

        // 3. 拨动第一棒
        chans[0] <- struct{}{}

        // 4. 等待
        wg.Wait()
        fmt.Println("\n任务完成")
}
func main() {
        buffer(3) // 你可以改成 5, 10 等任何正整数
        quit(4)
}