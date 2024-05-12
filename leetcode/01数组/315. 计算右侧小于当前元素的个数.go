package main

// 315. 计算右侧小于当前元素的个数

// 给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

// countSmaller .
// 同 LCR 170. 交易逆序对的总数
// 使用归并排序，在排序过程中统计每个元素右侧比其小的元素个数
// 因为在归并过程中比较两个数字的大小，同时还需要在最终结果result中确定位置，所以需要对排序的数字记录其在原数组的下标值
// 又由于可以直接通过下标在原数组中找到其对应的值，所以我么可以直接对下标数组进行归并排序即可
func countSmaller(nums []int) []int {
	result := make([]int, len(nums))

	// 构造一个下标数组并对其进行归并排序
	indexNums := make([]int, len(nums))
	for i := range nums {
		indexNums[i] = i
	}

	mergeSort315(nums, indexNums, result)

	return result
}

// mergeSort 归并排序，过程中统计右侧比其小的数字个数，返回排好序的下标数组
// nums：原数组
// indexNums：待排序的下标数组
// result：总结果
func mergeSort315(nums, indexNums []int, result []int) []int {
	// 若数组中只剩一个元素则直接返回
	if len(indexNums) <= 1 {
		return indexNums
	}

	// 每对将数组平均分成两个子数组，递归进行归并排序
	leftIndexNums := mergeSort315(nums, indexNums[:len(indexNums)/2], result)
	rightIndexNums := mergeSort315(nums, indexNums[len(indexNums)/2:], result)

	// 两个子数组排序完成后，合并两个子数组，同时统计右侧比其小的数字个数，
	// 同时按下标从小到大遍历两个子数组，数组内部已经是排好序的，按照大小依次将左右数组元素添加到结果集，同时统计右侧比其小的数字个数，
	// 若右侧元素先入结果集，则左侧剩余的所有元素都多了一个右侧比其小的数字，相反如果左边元素先插入，则说明不存在右侧比其小的数字
	resultIndexNums := make([]int, 0, len(indexNums))
	l, r := 0, 0
	for l < len(leftIndexNums) && r < len(rightIndexNums) {
		if nums[leftIndexNums[l]] <= nums[rightIndexNums[r]] {
			resultIndexNums = append(resultIndexNums, leftIndexNums[l])
			l++
		} else {
			resultIndexNums = append(resultIndexNums, rightIndexNums[r])
			r++
			// 若右侧元素先入结果集，说明左侧剩余元素都比其大且都在它左面，左侧剩余的所有元素都多了一个右侧比其小的数字
			for i := l; i < len(leftIndexNums); i++ {
				// 递增每一个左侧元素的右侧比其小的元素个数
				result[leftIndexNums[i]]++
			}
		}
	}

	// 上面遍历时任何一个子数组遍历完就会跳出循环，对剩下一个未遍历完的数组全部追加
	if l < len(leftIndexNums) {
		resultIndexNums = append(resultIndexNums, leftIndexNums[l:]...)
	} else {
		resultIndexNums = append(resultIndexNums, rightIndexNums[r:]...)
	}

	return resultIndexNums
}
