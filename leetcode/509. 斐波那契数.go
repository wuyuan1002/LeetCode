package main

// 509. 斐波那契数

func main() {

}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	fib2 := 0
	fib1 := 1
	res := 0
	for i := 2; i <= n; i++ {
		res = fib2 + fib1
		fib2 = fib1
		fib1 = res
	}
	return res
}
