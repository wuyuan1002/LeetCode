package main

import "container/list"

// 146. LRU缓存机制

// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

// LRUCache 最近最少使用机制
type LRUCache struct {
	// 容量
	cap int
	// 使用一个链表存储所有元素，链表头是最近访问过的元素，链表尾是最久未访问的元素，删除时从链表尾开始删除元素
	cacheList *list.List
	// 使用一个map存储所有元素，用于在O(1)复杂度内判断元素是否存在,以及在O(1)复杂度内获得到链表中的指定元素
	cacheMap map[int]*list.Element
}

type kv struct {
	k,
	v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:       capacity,
		cacheList: list.New(),
		cacheMap:  make(map[int]*list.Element),
	}
}

// Get .
func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cacheMap[key]; ok {
		// 若存在，将节点移动到链表头，后返回value
		lru.cacheList.MoveToFront(node)
		return node.Value.(kv).v
	}
	return -1
}

// Put .
func (lru *LRUCache) Put(key int, value int) {
	// 若存在，更新值，并移动到链表头
	if node, ok := lru.cacheMap[key]; ok {
		node.Value = kv{k: key, v: value}
		lru.cacheList.MoveToFront(node)
		return
	}

	// 若容量不够了，先删除最久未使用的节点
	if lru.cacheList.Len() == lru.cap {
		del := lru.cacheList.Back()
		delVal := lru.cacheList.Remove(del)
		delete(lru.cacheMap, delVal.(kv).k)
	}

	// 将新的元素添加到缓存中
	node := lru.cacheList.PushFront(kv{
		k: key,
		v: value,
	})
	lru.cacheMap[key] = node
}
