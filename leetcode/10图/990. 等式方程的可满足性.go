package main

// 990. 等式方程的可满足性

// 给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
//
// 只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。

// equationsPossible .
// 并查集
// 由于等式相等具有传递性，可以使用并查集进行相同元素的集合管理，将所有相等的元素属于同一个连通分量，
//
// 将 a==b 表示节点a与节点b之间存在边，a!=b 表示节点a与节点b之间不存在边
// 1. 遍历所有等式，构造并查集，连接所有边，所有相等的元素属于同一个连通分量
// 2. 遍历所有不等式，若存不等式中的两节点同属于一个连通分量则说明题目所给的所有等式不成立
func equationsPossible(equations []string) bool {
	// 构造并查集，初始时26个字母都是单独的连通分量
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	// 遍历所有等式 -- 构造并查集，连接所有边，所有相等的元素属于同一个连通分量
	for _, equation := range equations {
		if equation[1] == '=' {
			connect(parent, int(equation[0]-'a'), int(equation[3]-'a'))
		}
	}

	// 遍历所有不等式 -- 若存不等式中的两节点同属于一个连通分量则说明题目所给的所有等式不成立
	for _, equation := range equations {
		if equation[1] == '!' {
			if getRoot(parent, int(equation[0]-'a')) == getRoot(parent, int(equation[3]-'a')) {
				return false
			}
		}
	}

	return true
}
