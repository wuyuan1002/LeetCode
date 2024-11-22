package main

// 208. 实现 Trie (前缀树)

// Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。
// 这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

// Trie 描述一个前缀树 -- 一个26叉树
type Trie struct {
	isEndOfWord bool      // 标记当前节点是否为某一个单词的结尾节点
	children    [26]*Trie // 26个孩子节点，使用下标记录找到存储某个字符的下一个Trie
}

// Constructor 构造一个前缀树
func Constructor() Trie {
	return Trie{}
}

// Insert 插入一个单词
func (this *Trie) Insert(word string) {
	// 不断遍历要插入的单词的每个字符，将其加入到前缀树中
	node := this
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &Trie{}
		}
		node = node.children[c-'a']
	}
	// 将最后一个字符对应的节点标记为结尾字符
	node.isEndOfWord = true
}

// Search 在前缀树中查找给定单词，若能够找到该单词的路径并且最后一个字符对应的节点是结尾字符，则说明给定单词存在
func (this *Trie) Search(word string) bool {
	// 不断查找给定单词的每个字符是否存在前缀树中
	node := this
	for _, c := range word {
		// 若发现当前字符不存在则直接返回false
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}

	// 若找到了给定单词的完整路径，那么如果这个节点是结尾节点则说明该单词存在
	return node.isEndOfWord
}

// StartsWith 查找给定前缀是否存在前缀树中，与查找单词唯一的不同是：查找前缀不要求最后一个字符的节点是结尾节点
func (this *Trie) StartsWith(prefix string) bool {
	// 不断查找给定前缀的每个字符是否存在前缀树中
	node := this
	for _, c := range prefix {
		// 若发现当前字符不存在则直接返回false
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}

	// 最终发现给定前缀的所有字符都在前缀树中则说明存在此前缀，直接返回true
	return true
}
