package main

// 238. 除自身以外数组的乘积

// 给你一个长度为n的整数数组nums，其中n > 1，返回输出数组output，
// 其中 output[i]等于nums中除nums[i]之外其余各元素的乘积。

func main() {

}

// 同 offer 66
func productExceptSelf(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	result := make([]int, len(nums))
	for i := range result {
		result[i] = 1
	}

	// 乘nums里i的前半部分
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// 乘nums里i的后半部分
	accu := 1
	for i := len(nums) - 2; i >= 0; i-- {
		accu *= nums[i+1]
		result[i] *= accu
	}

	return result
}
