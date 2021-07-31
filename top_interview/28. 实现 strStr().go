package main

// 28. 实现 strStr()

// 实现strStr()函数。
// 给你两个字符串haystack 和 needle ，请你在 haystack 字符串中
// 找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
//
// 说明：
// 当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
// 对于本题而言，当needle是空字符串时我们应当返回 0 。
// 这与 C 语言的strstr()以及 Java 的indexOf()定义相符。

func main() {

}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	res := -1
	index := -1
	for i := 0; i < len(haystack); i++ {
		k, j := i, 0 // 一次比较时在两个字符串中的指针
		for ; k < len(haystack) && j < len(needle) && haystack[k] == needle[j]; k, j = k+1, j+1 {
			if j == 0 {
				// 记录起始位置
				index = i
			}
		}

		// 已找到
		if j == len(needle) {
			res = index
			break
		}
	}

	return res
}
