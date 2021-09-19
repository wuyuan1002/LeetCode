package main

// 650. 只有两个键的键盘

// 最初记事本上只有一个字符 'A' 。你每次可以对这个记事本进行两种操作：
//
// Copy All（复制全部）：复制这个记事本中的所有字符（不允许仅复制部分字符）。
// Paste（粘贴）：粘贴 上一次 复制的字符。
// 给你一个数字n ，你需要使用最少的操作次数，
// 在记事本上输出 恰好n个 'A' 。返回能够打印出n个 'A' 的最少操作次数。

func main() {

}

func minSteps(n int) int {
	sum, copyy, num := 1, 0, 0 // 当前记事本中A的个数、当前拷贝的所有A、累计操作的次数
	for sum < n {
		if n%sum == 0 {
			num++
			copyy = sum
		}

		num++
		sum += copyy
	}

	return num
}
