package main

// 208. 实现 Trie (前缀树)

// Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。
// 这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

// Trie .
type Trie struct {
	isEnd    bool      // 标记当前节点是否为结束节点 -- 当前节点是否为叶子节点
	children [26]*Trie // 26个孩子节点，使用下标记录找到存储某个字符的下一个Trie
}

// Constructor .
func Constructor() Trie {
	return Trie{}
}

// Insert .
func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &Trie{}
		}
		node = node.children[c-'a']
	}
	node.isEnd = true
}

// Search .
func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}

	return node.isEnd
}

// StartsWith .
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
