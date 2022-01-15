package main

// 剑指 Offer 40. 最小的k个数

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，
// 则最小的4个数字是1、2、3、4。

// func main() {
// 	// 最小的k个数、最小的第k个数、最大的k个数、最大的第k个数 -- 使用同样的方法
// }

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	start, end := 0, len(arr)-1
	index := partition(arr, start, end)
	for index != k-1 { // 一旦index==k-1，即可停止。
		if index > k-1 {
			end = index - 1
		} else {
			start = index + 1
		}
		index = partition(arr, start, end)
	}
	return arr[:k] // 返回的这k个元素不一定是有序的
}

func partition(arr []int, left, right int) int {

	l, r := left, right
	temp := arr[left]
	for l < r {
		for arr[r] >= temp && l < r {
			r--
		}
		for arr[l] <= temp && l < r {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[left], arr[l] = arr[l], temp
	return l
}

// 求最小的第k个数
func getLeastNumbersk(arr []int, k int) int {
	if arr == nil || k <= 0 {
		panic("")
	}
	if k >= len(arr) {
		panic("")
	}

	start, end := 0, len(arr)-1
	index := partition(arr, start, end)
	for index != k-1 {
		if index < k-1 {
			start = index + 1
		} else {
			end = index - 1
		}
		index = partition(arr, start, end)
	}
	return arr[k-1]
}
