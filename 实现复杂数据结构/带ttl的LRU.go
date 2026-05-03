/*
思路概括：
1. 在标准 LRU 基础上为每个节点添加过期时间 expireAt（写入时计算绝对过期时间）。
2. 访问时（Get/getNode）先检查是否过期：若过期则被动删除该节点并返回未命中。
3. Put 时创建新节点并计算绝对过期时间，容量超标时淘汰最久未用的尾节点。
4. 当前实现仅适用于单线程环境，高并发场景需分段锁、sync.Pool 等优化。

复杂度：
- 时间：Get / Put 均为 O(1)。
- 空间：O(capacity)。
*/

//当前版本只是演示
//在高并发环境下会出现线程安全问题
//原因是map不是并发安全的，不能并发读写
//此外如果不加锁，在并发环境下，修改双向链表指针也有可能使得指针指向错误的地址
//同时，淘汰逻辑只有被动淘，高并发场景下可能会导致内存占用虚高

//最合适的解决方法是分段锁，把大的LRU缓存拆分成多个小弟lru实例
//利用hash分片将请求路由到对应分片
//此外还可以引入sync.Pool建立节点池，避免内存分配和GC的开销

package main

import (
	"fmt"
	"time"
)

//题目实现前必须要想清楚怎么写数据结构和对应的方法
//首先要实现Node和LRUCache两个结构体

//node基础包含键值对，前后指针
//如果是带ttl模式需要加上过期时间
type Node struct {
	key, value int
	pre, next  *Node
	expireAt   time.Time // 绝对过期时间
}

//LRU缓存要包括容量，链表，map
//如果是带TTl模式需要加上ttl
type LRUCache struct {
	capacity  int
	ttl       time.Duration
	dummy     *Node
	keyToNode map[int]*Node
}

//先写构造函数，构造函数唯一要点是必须把dummy节点的前后指针指向自身
//这样不用进行边界判断，也即无论链表有多长，dummy的前后都有节点，即使是自己
func Constructor(capacity int, ttl time.Duration) LRUCache {
	dummy := &Node{}
	dummy.next = dummy
	dummy.pre = dummy
	return LRUCache{
		capacity:  capacity,
		ttl:       ttl,
		dummy:     dummy,
		keyToNode: map[int]*Node{},
	}
}

//再写删除node，压入node，获取node置三个操作
//注意这里直接把更新node的位置封装在了getnode内部
func (l *LRUCache) removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (l *LRUCache) pushNode(node *Node) {
	node.pre = l.dummy
	node.next = l.dummy.next
	l.dummy.next = node
	node.next.pre = node
}

func (l *LRUCache) getNode(key int) *Node {
	node := l.keyToNode[key]
	if node == nil {
		return nil
	}
	// 被动删除：访问时发现已过期
	if time.Now().After(node.expireAt) {
		delete(l.keyToNode, node.key)
		l.removeNode(node)
		return nil
	}
	l.removeNode(node)
	l.pushNode(node)
	return node
}





//再写暴露在外的接口Get和Put
func (l *LRUCache) Get(key int) int {
	node := l.getNode(key)
	if node == nil {
		return -1
	}
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	oldNode := l.getNode(key)
	if oldNode != nil {
		oldNode.value = value
		return
	}
	node := &Node{
		key:      key,
		value:    value,
		expireAt: time.Now().Add(l.ttl), // 直接算出绝对过期时间
	}
	l.keyToNode[key] = node
	l.pushNode(node)
	if len(l.keyToNode) > l.capacity {
		last := l.dummy.pre
		delete(l.keyToNode, last.key)
		l.removeNode(last)
	}
}

func main() {
	cache := Constructor(10, 2*time.Second)

	cache.Put(1, 1314)
	fmt.Println("写入后立即读:", cache.Get(1)) // 1314

	time.Sleep(1 * time.Second)
	fmt.Println("1秒后读:", cache.Get(1)) // 1314

	time.Sleep(2500 * time.Millisecond)
	fmt.Println("再等2.5秒后读:", cache.Get(1)) // -1
}
