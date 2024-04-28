package main

// 剑指 Offer 51. 数组中的逆序对

// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
// 输入一个数组，求出这个数组中的逆序对的总数。

// func main() {

// }

func reversePairs(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	_, n := sort(nums)
	return n
}

// @param nums 待排序的数组
// @return []int 排好序的数组
// @return int 逆序对个数
func sort(nums []int) ([]int, int) {
	// 递归结束条件：数组中只剩下一个元素
	if len(nums) <= 1 {
		return nums, 0
	}

	// 每对将数组平均分成两个子数组，作递归
	left, n1 := sort(nums[:len(nums)/2])
	right, n2 := sort(nums[len(nums)/2:])

	// 从小到大排序，合并排好序的两个子数组, 逻辑为：
	// 同时按下标从小到大遍历两个子数组，数组内部已经是排好序，只需要比较两个子数组之间的元素
	// 逆序对就在此时计算出来，在组建结果集的过程中，
	// 如果右侧元素先入结果集，则左侧剩余的所有元素都与它组成逆序对，相反如果左边元素先插入，则说明不存在逆序对
	res := make([]int, 0, len(left)+len(right))
	i, j, reverseNum := 0, 0, n1+n2
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
			reverseNum += len(left) - i
		}
	}

	// 上面遍历时任何一个子数组遍历完就会跳出循环，对剩下一个未遍历完的数组全部追加
	if i < len(left) {
		res = append(res, left[i:]...)
	} else {
		res = append(res, right[j:]...)
	}

	return res, reverseNum
}

// 第二次做 -- 归并排序，从小到大排
// @param nums 待排序的数组
// @return []int 排好序的数组
// @return int 逆序对个数
func sort1(nums []int) ([]int, int) {
	// 数组拆分的最小数组
	if len(nums) <= 1 {
		return nums, 0
	}

	// 将数组拆分，排序左右两个字数组
	left, nl := sort1(nums[:len(nums)/2])
	right, nr := sort1(nums[len(nums)/2:])

	// 将两个排好序的子数组从小到大合并
	res := make([]int, 0, len(left)+len(right))
	l, r, reverseNum := 0, 0, nl+nr
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] { // 左面数字比右面数字小，表明前面数字比后面数字小
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
			reverseNum += len(left) - l
		}
	}

	if l < len(left) {
		res = append(res, left[l:]...)
	}
	if r < len(right) {
		res = append(res, right[r:]...)
	}

	return res, reverseNum
}
