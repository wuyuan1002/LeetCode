package main

// LCR 170. 交易逆序对的总数

// 在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。
// 请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的「交易逆序对」总数。

// reversePairs .
// 同 Offer 51. 数组中的逆序对
// 归并排序，过程中统计逆序对个数
func reversePairs(record []int) int {
	_, reverseNum := mergeSort(record)
	return reverseNum
}

// mergeSort 归并排序，过程中统计逆序对个数
// nums：待排序的数组
// 返回：排好序的数组、该数组中逆序对个数
func mergeSort(nums []int) ([]int, int) {
	// 递归结束条件：数组中只剩下一个元素
	if len(nums) <= 1 {
		return nums, 0
	}

	// 每对将数组平均分成两个子数组，递归进行归并排序
	left, leftReverseNum := mergeSort(nums[:len(nums)/2])
	right, rightReverseNum := mergeSort(nums[len(nums)/2:])

	// 两个子数组排序完成后，合并两个子数组，同时统计逆序对个数，逻辑为：
	// 同时按下标从小到大遍历两个子数组，数组内部已经是排好序，只需要比较两个子数组之间的元素
	// 逆序对就在此时计算出来，在组建结果集的过程中，
	// 如果右侧元素先入结果集，则左侧剩余的所有元素都与它组成逆序对，相反如果左边元素先插入，则说明不存在逆序对
	resultNums := make([]int, 0, len(left)+len(right))
	i, j, reverseNum := 0, 0, leftReverseNum+rightReverseNum
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			resultNums = append(resultNums, left[i])
			i++
		} else {
			resultNums = append(resultNums, right[j])
			j++
			reverseNum += len(left) - i
		}
	}

	// 上面遍历时任何一个子数组遍历完就会跳出循环，对剩下一个未遍历完的数组全部追加
	if i < len(left) {
		resultNums = append(resultNums, left[i:]...)
	} else {
		resultNums = append(resultNums, right[j:]...)
	}

	return resultNums, reverseNum
}
