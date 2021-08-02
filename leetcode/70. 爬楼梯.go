package main

// 70. 爬楼梯

func main() {

}

// offer 10-2
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	c2 := 1
	c1 := 1
	res := 0
	for i := 2; i <= n; i++ {
		res = c1 + c2
		c2 = c1
		c1 = res
	}
	return res
}
