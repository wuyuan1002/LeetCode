package main

// 剑指 Offer 62. 圆圈中最后剩下的数字

// 0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。
// 求出这个圆圈里剩下的最后一个数字。
//
// 例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，
// 则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。

// func main() {

// }

// 方法1：用环形链表或数组模拟圆圈 -- 但真正模拟删除操作开销比较大，数组每次都需要把后面的元素向前移，链表每次都需要寻找后面第m个要删除的元素
// 方法2：约瑟夫环
func lastRemaining(n int, m int) int {
	if n < 1 || m < 1 {
		return -1
	}

	start := 0
	for i := 2; i <= n; i++ {
		start = (start + m) % i
	}
	return start
}
