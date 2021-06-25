package main

// 647. 回文子串

// 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

func main() {

}

// Manacher算法
func countSubstrings(s string) int {
	n := len(s)
	t := "$#"
	for i := 0; i < n; i++ {
		t += string(s[i]) + "#"
	}
	n = len(t)
	t += "!"

	f := make([]int, n)
	iMax, rMax, ans := 0, 0, 0
	for i := 1; i < n; i++ {
		// 初始化 f[i]
		if i <= rMax {
			f[i] = min(rMax-i+1, f[2*iMax-i])
		} else {
			f[i] = 1
		}
		// 中心拓展
		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}
		// 动态维护 iMax 和 rMax
		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}
		// 统计答案, 当前贡献为 (f[i] - 1) / 2 上取整
		ans += f[i] / 2
	}
	return ans
}
