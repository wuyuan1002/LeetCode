package main

// 376. 摆动序列

// 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。
//
// 例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
//
// 相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
// 子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
//
// 给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。

// wiggleMaxLength .
// 同 leetcode 1567. 乘积为正数的最长子数组长度
//
// up[i]表示以第i个数字结尾的最后两个数字递增的最长摆动子序列长度 -- 如[1, 6, 3, 8, 5, 7]为递增子序列（7比5大，最后两个数字递增）
// down[i]表示以第i个数字结尾的最后两个数字递减的最长摆动子序列长度 -- 如[1, 6, 3, 8, 5]为递减子序列（5比8小，最后两个数字递减）
func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 构造两个dp数组，分别表示递增摆动序列长度和递减摆动序列长度
	up, down := make([]int, len(nums)), make([]int, len(nums))

	// 初始化dp数组 -- 数组中第一个数字即是递增摆动子序列也是递减摆动子序列，长度为1
	up[0], down[0] = 1, 1

	// 从第二个数字开始遍历每一个数字，分别计算在[0, i]区间内最长递增摆动子序列和递减子序列的长度
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			// 若当前数字大于前一个数字，以当前数字结尾所构成的摆动序列一定是递增序列，
			// 更新当前下标能构成的递增摆动子序列的最大长度，且当前下标内递减摆动序列长度为前一个下标内的递减摆动序列长度不变
			// 递增序列的前一个一定是递减序列，也就是说递增序列一定是由递减序列产生
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			// 若当前数字小于前一个数字，以当前数字结尾所构成的摆动序列一定是递减序列
			// 更新当前下标能构成的递减摆动子序列的最大长度，且当前下标内递增摆动序列长度为前一个下标内的递增摆动序列长度不变
			// 递减序列的前一个一定是递增序列，也就是说递减序列一定是由递增序列产生
			down[i] = max(up[i-1]+1, down[i-1])
			up[i] = up[i-1]
		} else {
			// 若当前数字与前一个数字相同，说明当前数字所能构成的递增和递减序列长度与前一个数字的一样保持不变
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}

	// 返回最终数组中递增摆动子序列和递减摆动子序列长度的最大值
	return max(up[len(nums)-1], down[len(nums)-1])
}

// wiggleMaxLength1 滚动数组优化
// 可以看出在求up[i]和down[i]时，只依赖其前一个数字的结果，即up[i-1]和down[i-1]，
// 因此，up和down可以只使用一个变量用来记录前一个下标所能构成的最长递增摆动子序列和最长递减摆动子序列的长度即可，
// 而不必使用数组将每一个下标的都记录下来
func wiggleMaxLength1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 使用两个变量记录数组在[0, i]范围内所能构成的最长递增摆动子序列和最长递减摆动子序列的长度
	// 数组只有的第一个元素即构成递增摆动序列也构成递减摆动序列，所以初始值为1
	up, down := 1, 1

	// 从第二个数字开始遍历每一个数字，分别计算在[0, i]区间内最长递增摆动子序列和递减子序列的长度
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			// 若当前数字大于前一个数字，以当前数字结尾所构成的摆动序列一定是递增序列
			// 递增序列一定由递减序列产生，且此时所记录的递减序列长度一定是最大值
			up = down + 1
		} else if nums[i] < nums[i-1] {
			// 若当前数字大于前一个数字，以当前数字结尾所构成的摆动序列一定是递减序列
			// 递减序列一定由递增序列产生，且此时所记录的递增序列长度一定是最大值
			down = up + 1
		}
		// 若两数字相同，说明[0, i]范围内所能构成的递增和递减序列长度都不变
	}

	// 返回递增摆动子序列长度和递减子序列长度的最大值
	return max(up, down)
}
