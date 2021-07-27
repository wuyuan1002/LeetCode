package main

import (
	"container/list"
)

// 460. LFU 缓存

// 请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
//
// 实现 LFUCache 类：
//
// LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
// int get(int key) - 如果键存在于缓存中，则获取键的值，否则返回 -1。
// void put(int key, int value) - 如果键已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
// 注意「项的使用次数」就是自插入该项以来对其调用 get 和 put 函数的次数之和。使用次数会在对应项被移除后置为 0 。
//
// 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
//
// 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

func main() {
	obj := Constructor(0)
	obj.Put(2, 22)
	obj.Get(0)
	// obj.Put(3, 33)
	// obj.Put(4, 44)
	// obj.Put(3, 33)
	// obj.Put(5, 55)
}

// LFUCache 类似 146. LRU缓存机制
type LFUCache struct {
	cap       int
	cacheList *list.List // 按照使用次数由大到小排序，相同使用次数的，将最新的放到前面，保证最新使用的在链表前面，这样链表最后的就总是最近要被删除的
	cacheMap  map[int]*list.Element
}

type kv struct {
	count int // 使用次数
	k, v  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:       capacity,
		cacheList: list.New(),
		cacheMap:  make(map[int]*list.Element),
	}
}

func (c *LFUCache) Get(key int) int {
	if node, ok := c.cacheMap[key]; ok {
		// 若存在，将节点使用次数+1，并将节点移动到相同使用次数节点的第一个位置

		// 1. 增加使用次数
		count := node.Value.(kv).count
		node.Value = kv{
			count: count + 1,
			k:     node.Value.(kv).k,
			v:     node.Value.(kv).v,
		}

		// 2. 将节点移动到相同使用次数节点的第一个位置
		c.move(node)

		return node.Value.(kv).v
	}

	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if c.cap <= 0 {
		return
	}

	if node, ok := c.cacheMap[key]; ok {
		// 若存在，更新值，移动节点位置
		count := node.Value.(kv).count
		node.Value = kv{
			count: count + 1,
			k:     key,
			v:     value,
		}
		c.move(node)
		return
	}

	// 若不存在，判断容量，加入新的节点
	if c.cacheList.Len() >= c.cap {
		// 若容量不够了，先删掉链表尾的一个元素
		del := c.cacheList.Back()
		c.cacheList.Remove(del)
		delete(c.cacheMap, del.Value.(kv).k)
	}

	// 添加新的节点 -- 先加到链表尾，再向前移动
	node := c.cacheList.PushBack(kv{
		count: 1,
		k:     key,
		v:     value,
	})
	c.move(node)
	c.cacheMap[key] = node
}

// 将节点移动到相同使用次数节点的最前面位置
func (c *LFUCache) move(node *list.Element) {
	pre := node.Prev()
	if pre == nil {
		return
	}
	for ; pre != nil && pre.Value.(kv).count <= node.Value.(kv).count; pre = pre.Prev() {
	}

	if pre != nil {
		c.cacheList.MoveAfter(node, pre)
	} else {
		c.cacheList.MoveToFront(node)
	}
}
