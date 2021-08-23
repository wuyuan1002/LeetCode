package main

// 198. 打家劫舍

func main() {

}

func rob1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	pre2, pre1 := nums[0], max(nums[0], nums[1]) // 前两个和前一个
	for i := 2; i < len(nums); i++ {
		price := max(pre2+nums[i], pre1)
		pre2 = pre1
		pre1 = price
	}

	return pre1
}
