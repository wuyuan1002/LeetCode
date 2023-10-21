package main

// 496. 下一个更大元素 I

// 给你两个 没有重复元素 的数组nums1 和nums2，其中nums1是nums2的子集。
// 请你找出 nums1中每个元素在nums2中的下一个比其大的值。
// nums1中数字x的下一个更大元素是指x在nums2中对应位置的右边的第一个比x大的元素。
// 如果不存在，对应位置输出 -1 。

// nextGreaterElement .
// 单调栈 -- 单调递增栈
// 求一个数组中每个数字的下一个比其大的元素 -- 使用单调递增栈，栈内元素都是还没出现比其大的元素的元素（等待被计算的元素）
// 不断从前到后遍历数组，当遍历到的当前元素比栈顶元素小时，说明比栈内元素和当前元素大的元素还没出现，直接将当前元素入栈等待，
// 当遍历到的当前元素比栈顶元素大时，说明当前元素是栈内比其小的元素的下一个更大元素，开始将栈内元素挨个出栈，进行计算，
// 直到栈内元素都比当前元素大时，说明比栈内剩余元素和当前元素大的元素还没出现，将当前元素入栈等待，继续向后遍历数组
// 当数组全部遍历完成后，栈内元素就都是没有找到其右边比其大的元素了
//
// 本题求的是数组中每个元素的下一个更大元素的值
// 本题是求nums2中一部分元素（nums1是nums2的子集）的下一个比其大的元素，
// 我们可以求出nums2数组中每个元素右边比其大的元素的值，先全部使用map存起来，
// 然后根据nums1中的数字在全部数字中选取
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int, len(nums2)) // 定义一个map用来存nums2中每个元素和其下一个更大元素的值
	stack := make([]int, 0)               // 单调递增栈，用来存还没有出现下一个比其大的元素的值

	// 从头到尾遍历nums2，先全量求出nums2中每个数字的下一个比其大的元素，使用一个map进行存储
	// 遍历结束后，map中存着的就是存在下一个比其大的元素的值和更大元素的值，
	// map中没存的，就是在单调栈中的元素，也就是在它们之后没有出现比其更大的元素了
	for i := 0; i < len(nums2); i++ {
		// 当当前元素比栈顶元素大时，说明当前元素是栈内部分元素的下一个更大元素
		// 将栈内元素挨个出栈，记录栈内比当前元素小的元素的下一个更大元素为当前元素
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			// 获取栈顶元素并出栈
			num := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 记录栈顶元素的下一个更大元素为当前元素，并保存到map中
			hash[num] = nums2[i]
		}

		// 栈内剩下的元素都是大于等于当前元素的元素了
		// 将当前元素入栈，等待自己的下一个更大元素出现后记录到map中
		stack = append(stack, nums2[i])
	}

	// 遍历nums1，从map中获取其每个数字的下一个更大元素的值，若map中没有，说明其后面没有下一个更大元素
	result := make([]int, len(nums1))
	for i, num := range nums1 {
		if maxNum, ok := hash[num]; ok {
			result[i] = maxNum
		} else {
			result[i] = -1
		}
	}
	return result
}
