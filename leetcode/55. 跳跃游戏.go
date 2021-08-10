package main

// 55. 跳跃游戏

func main() {

}

// 从某一位置起跳，能跳到的最远位置之前的位置都是可以跳到的。
// 所以，遍历每一个位置，不断更新最远位置，最远位置之前的都是可以跳到的
func canJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return true
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	far := 0 // 能跳到的最远位置
	for i := 0; i < len(nums); i++ {
		if i > far {
			return false
		}
		far = max(far, i+nums[i])
	}

	return true
}
