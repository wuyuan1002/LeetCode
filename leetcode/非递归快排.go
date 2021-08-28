package main

// 非递归快排

func main() {

}

func quickSort1(nums []int, left, right int) {
	stack := []int{left, right} // 使用栈保存每次需要递归的左右界限
	for len(stack) > 0 {
		// 获取本次partition的左右界限
		right = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 开始partition
		l, r := left, right
		tmp := nums[left]
		for l < r {
			for l < r && nums[r] >= tmp {
				r--
			}
			for l < r && nums[l] <= tmp {
				l++
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
		}
		nums[l], nums[left] = tmp, nums[l]

		// 将下一次需要partition的左右边界入栈
		stack = append(stack, left, l)
		stack = append(stack, r+1, right)
	}
}
