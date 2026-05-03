/*
思路概括：
1. 使用双向链表存储键值对节点，链表头为最近使用，尾为最久未用。
2. 通过哈希表 keyToNode 实现 O(1) 时间定位节点。
3. Get 时：命中则将节点移至链表头部（removeNode + pushNode），未命中返回 -1。
4. Put 时：若 key 已存在则更新值并移到头部；否则新建节点插入头部，容量超标时淘汰尾部节点。

复杂度：
- 时间：Get / Put 均为 O(1)。
- 空间：O(capacity)，即哈希表和链表节点所占空间。
*/

// 背诵技巧，先定义Node和LRUCache结构体，
// node是一个双向链表节点
// LRUCache需要包含容量，双向链表头节点，hash表
// 先实现三个中间层函数，获取，插入，删除node
// 再实现暴露在外的接口Get和Put，Get只需要接收键返回值，Put只需要插入键值对
// 暴露的api本质上只接受键值对，内部实现链表
package main

import "fmt"

type Node struct {
	key, value int
	pre, next  *Node
}
type LRUCache struct {
	capacity  int
	dummy     *Node
	keyToNode map[int]*Node
}

//新建结构体函数+
//增删查改四个函数

func Constructor(capacity int) LRUCache {
	dummy := &Node{}
	dummy.next = dummy
	dummy.pre = dummy
	return LRUCache{
		capacity:  capacity,
		dummy:     dummy,
		keyToNode: map[int]*Node{},
	}
}

//增加函数暴露在外的接口是Put，
//需要实现在map和list中保留对应元素,对应压入
// 容量超标时，还需要实现在map和list中删掉键值对和node

func (l *LRUCache) removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
	//不需要修改剩下的两个指针，
	//原因是go语言自动垃圾回收，扫描发现node没有被引用就会将其回收
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
	l.removeNode(node)
	l.pushNode(node)
	return node

}

func (l *LRUCache) Put(key int, value int) {
	oldNode := l.getNode(key)
	if oldNode != nil { //已经有键相同的node了，只需要更新值
		oldNode.value = value
		return
	}
	//需要插入新的Node
	node := &Node{
		key:   key,
		value: value,
	}
	l.keyToNode[key] = node
	l.pushNode(node)
	//如果超出容量，必须删除最后一个node
	if len(l.keyToNode) > l.capacity {
		lastNode := l.dummy.pre
		delete(l.keyToNode, lastNode.key)
		l.removeNode(lastNode)
	}

}

func (l *LRUCache) Get(key int) int {
	node := l.getNode(key)
	if node == nil {
		return -1
	}
	return node.value
}

func main() {
	testLru := Constructor(10)
	testLru.Put(1, 1314)
	ans := testLru.Get(1)
	fmt.Println(ans)
}
