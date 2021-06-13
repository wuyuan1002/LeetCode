package main

import (
	"container/list"
)

// 146. LRU 缓存机制

// 运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
// 实现 LRUCache 类：
//
// LRUCache(int capacity) 以正整数作为容量capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value)如果关键字已经存在，则变更其数据值；
// 如果关键字不存在，则插入该组「关键字-值」。
// 当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
// 进阶：你是否可以在O(1) 时间复杂度内完成这两种操作？

func main() {

}

// LRUCache 最近最少使用机制
type LRUCache struct {
	// 记录缓存的容量
	cap int
	// 使用一个链表存储所有元素，链表头是最近访问过的元素，链表尾是最久未访问的元素，删除时从链表尾开始删除元素
	cacheList *list.List
	// 使用一个map存储所有元素，用于在O(1)复杂度内判断元素是否存在,以及在O(1)复杂度内获得到链表中的指定元素
	cacheMap map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:       capacity,
		cacheList: list.New(),
		cacheMap:  make(map[int]*list.Element),
	}
}

type KV struct {
	K, V int
}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.cacheMap[key]; ok {
		// 若存在，将节点移动到链表头，后返回value
		cache.cacheList.MoveToFront(node)
		return node.Value.(KV).V
	}
	// 若不存在则返回-1
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	// 若存在，更新值，并移动到链表头
	if node, ok := cache.cacheMap[key]; ok {
		node.Value = KV{key, value}
		cache.cacheList.MoveToFront(node)
		return
	}

	// 若不存在，容量够，则链表头新增元素，容量不够，先删除链表尾元素，后在链表头新增元素
	if cache.cacheList.Len() < cache.cap {
		// 添加到链表头
		node := cache.cacheList.PushFront(KV{key, value})
		// 添加到map
		cache.cacheMap[key] = node
	} else {
		// 获取最后一个节点
		lastNode := cache.cacheList.Back()
		// zaimap中删除
		delete(cache.cacheMap, lastNode.Value.(KV).K)
		// 在链表中删除
		cache.cacheList.Remove(lastNode)

		// 添加新节点到链表头
		node := cache.cacheList.PushFront(KV{key, value})
		// 添加到map
		cache.cacheMap[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
