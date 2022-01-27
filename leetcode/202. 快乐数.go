package main

// 202. 快乐数

// 编写一个算法来判断一个数 n 是不是快乐数。
//
// 「快乐数」定义为：
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果 可以变为 1，那么这个数就是快乐数。
// 如果 n 是快乐数就返回 true ；不是，则返回 false 。

// func main() {

// }

// 使用map存某依次的计算结果是否出现过，若出现过，则说明将会无限循环永远变不成1
func isHappy(n int) bool {
	mapp := make(map[int]struct{})
	getSum := func(num int) int { // 获取num每一位的平方和
		sum := 0
		for num != 0 {
			n := num % 10
			num /= 10
			sum += n * n
		}
		return sum
	}

	num := n
	for {
		num = getSum(num)
		if num == 1 {
			return true
		} else if _, ok := mapp[num]; ok {
			return false
		}
		mapp[num] = struct{}{}
	}
}

// 第二次做
func isHappy1(n int) bool {

	// 获取每一位数字的平方和
	getSum := func(num int) int {
		sum := 0
		for num != 0 {
			n := num % 10
			num /= 10
			sum += n * n
		}
		return sum
	}

	// 用于存放每次的平方和结果
	hash := make(map[int]bool)
	num := n
	for {
		num = getSum(num)
		if num == 1 {
			return true
		} else if _, ok := hash[num]; ok {
			// 次次结果之前出现过，说明出现了循环，不可能是快乐数
			return false
		} else {
			// 若没出现过，将结果记录下来，进入下一次计算
			hash[num] = true
		}
	}
}
