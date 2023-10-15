package main

// 131. 分割回文串

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。

// partition .
func partition(s string) [][]string {
	res := make([]string, 0)      // 存放一次回溯路径中每一层所选的满足条件的回文字符串
	result := make([][]string, 0) // 总结果集
	dfsPartition(s, 0, &res, &result)
	return result
}

// dfsPartition .
// s: 选择列表
// start: 每次遍历的起始下标 -- 指定当前层的选择范围 -- s[start: ]
// res: 一次回溯过程中的结果, 里面存放一次回溯路径中每一层所选的满足条件的回文字符串 -- 回溯路径
// result: 总结果集
func dfsPartition(s string, start int, res *[]string, result *[][]string) {
	// 若本层的开始位置等于字符串整体长度, 说明已经遍历完了整个字符串
	// 回溯路径中的结果就是一次递归路径的满足条件的结果
	// 将本次回溯路径结果记入总结果集
	if start == len(s) {
		tmp := make([]string, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	// 本层选择范围为 [start: len(s)] -- 先在本层固定一个回文串, 然后进入下一层固定下一个回文串, 直到整个字符串寻找完成即完成了一次回溯找到了一个结果
	// 在本层从start位置开始向后寻找回文串, 找到一个就放入集合进入下一层递归（每一层找到的回文串都是从短到长的）,
	// 开始从下一层选择回文串直到整个字符串寻找完成, 此时就完成了一次寻找 -- 将回溯路径中的结果记入总结果集
	// 然后回溯到上一层, 在上一层寻找新的回文串, 找到后再次进入下一层进行递归
	for i := start; i < len(s); i++ {
		// 不断在本层寻找回文串加入回溯路径, 然后进入下一层在剩余字符中寻找下一层的回文串
		if isPartion(s, start, i) {
			// 将本层当前选好的回文串加入回溯路径
			*res = append(*res, s[start:i+1])
			// 进入下一层在剩余字符中寻找下一层的回文串
			dfsPartition(s, i+1, res, result)
			// 下一层的回文串已经找完了, 将本层该回文串移出回溯路径, 继续在本层寻找下一个回文串
			*res = (*res)[:len(*res)-1]
		}
	}
}

// isPartion 判断是否为回文串
func isPartion(s string, l, r int) bool {
	for ; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
