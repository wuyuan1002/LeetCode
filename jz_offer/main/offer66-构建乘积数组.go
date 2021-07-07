package main

import (
	"fmt"
)

// 剑指 Offer 66. 构建乘积数组

// 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，
// 其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积,
// 即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}

func constructArr(a []int) []int {
	if a == nil || len(a) == 0 {
		return nil
	}

	result := make([]int, len(a))
	for i := range result {
		result[i] = 1
	}

	// 乘nums里i的前半部分
	for i := 1; i < len(a); i++ {
		result[i] = result[i-1] * a[i-1]
	}

	// 乘nums里i的后半部分
	accu := 1
	for i := len(a) - 2; i >= 0; i-- {
		accu *= a[i+1]
		result[i] *= accu
	}

	return result
}
