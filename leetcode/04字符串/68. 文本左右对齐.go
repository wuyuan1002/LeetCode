package main

import "strings"

// 68. 文本左右对齐

// 给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
//
// 你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
//
// 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
//
// 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
//
// 注意:
//
// 单词是指由非空格字符组成的字符序列。
// 每个单词的长度大于 0，小于等于 maxWidth。
// 输入单词数组 words 至少包含一个单词。

// fullJustify .
// + 对于每一行，首先确定最多可以放置多少个单词numWords，从而可以得到该行的空格个数numSpaces
// + 接下来有三种情况：
//   - 当前行是最后一行：单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格
//   - 当前行不是最后一行，且只有一个单词：该单词左对齐，在行末填充空格
//   - 当前行不是最后一行，且不只一个单词：
//     需要将空格均匀分配在单词之间，则每个单词之间应至少有 avgSpaces = numSpaces / (numWords-1) 个空格，
//     剩余 extraSpaces = numSpaces % (numWords−1) 个空格应该填充在前extraSpaces个单词之间，
//     因此，前extraSpaces个单词之间填充 avgSpaces+1 个空格，其余单词之间填充 avgSpaces 个空格
func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)

	left, right := 0, 0 // 每一行的第一个单词下标和最后一个单词下标 -- [left, right)
	for {
		left = right // 更新当前行的第一个单词下标为前一行的最后一个单词下标后一位

		// 循环确定当前行可以放多少单词，注意单词之间应至少有一个空格
		sumLen := 0 // 统计当前行单词长度之和 -- 不包含中间的空格
		for right < len(words) && sumLen+len(words[right])+right-left <= maxWidth {
			sumLen += len(words[right])
			right++
		}

		// 若当前行是最后一行 -- 单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格
		if right == len(words) {
			s := strings.Join(words[left:], " ")
			result = append(result, s+strings.Repeat(" ", maxWidth-len(s)))
			return result
		}

		// 确定当前行所放的单词个数和空格个数
		numWords, numSpaces := right-left, maxWidth-sumLen

		// 若当前行只有一个单词 -- 该单词左对齐，在行末填充剩余空格
		if numWords == 1 {
			result = append(result, words[left]+strings.Repeat(" ", numSpaces))
			continue
		}

		// 若当前行不只一个单词 -- 将空格平均放置在每个单词之间，并将多余的空格分配给前面的单词之间
		avgSpaces, extraSpaces := numSpaces/(numWords-1), numSpaces%(numWords-1)
		s := strings.Builder{}
		for i := 0; i < numWords; i++ {
			// 填充单词
			s.WriteString(words[left+i])
			// 若当前不是最后一个单词则在单词后面添加空格
			if i < numWords-1 {
				// 添加avgSpaces个空格
				s.WriteString(strings.Repeat(" ", avgSpaces))
				if extraSpaces > 0 {
					// 若还有多余的空格，则在当前间隙多加一个空格
					s.WriteString(" ")
					extraSpaces--
				}
			}
		}
		result = append(result, s.String())
	}
}
