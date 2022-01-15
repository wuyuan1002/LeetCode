package main

import (
	"strconv"
	"strings"
)

// 297. 二叉树的序列化与反序列化

// 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
// 同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
//
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
// 你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
//
// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅LeetCode 序列化二叉树的格式。
// 你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

// func main() {

// }

// 同offer 37

type Codec struct {
	// 编码时用
	str strings.Builder

	// 解码时用
	nodes []string
	index int
}

func Constructor3() Codec {
	return Codec{
		str: strings.Builder{},

		nodes: make([]string, 0),
		index: 0,
	}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		c.str.WriteString("nil,")
		return c.str.String()
	}

	c.str.WriteString(strconv.Itoa(root.Val))
	c.str.WriteString(",")

	_ = c.serialize(root.Left)
	_ = c.serialize(root.Right)

	return c.str.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	c.nodes = strings.Split(data, ",")
	c.index = 0

	return c.deser()
}

func (c *Codec) deser() *TreeNode {
	if c.nodes[c.index] == "nil" {
		c.index++
		return nil
	}

	val, _ := strconv.Atoi(c.nodes[c.index])
	c.index++
	root := &TreeNode{
		Val:   val,
		Left:  c.deser(),
		Right: c.deser(),
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
