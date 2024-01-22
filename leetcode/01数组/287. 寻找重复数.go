package main

// 287. 寻找重复数

// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），
// 可知至少存在一个重复的整数。
//
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
//
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

// findDuplicate .
// 若数组排序，则可以从后向前遍历数组或者使用二分法，来寻找第一个与下标不相等的数字
// 但是现在数组不排序，由于数组中数字都在[1, n]范围内，所以数组下标和数组中数字的映射关系可以抽象成一个链表，
// 数组中如果有重复的数，那么就会产生多对一的映射，这样，形成的链表就一定会有环路了
//
// 综上
// 1.数组中有一个重复的整数 -> 链表中存在环
// 2.找到数组中的重复整数 -> 找到链表的环入口
// 同 leetcode 141. 环形链表 的解决思路，快慢指针
func findDuplicate(nums []int) int {
	slow, fast := 0, 0 // 快慢指针
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			fast = 0
			for slow != fast {
				slow = nums[slow]
				fast = nums[fast]
			}
			return slow
		}
	}
}
