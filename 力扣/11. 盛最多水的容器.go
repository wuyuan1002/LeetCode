package main

// 11. 盛最多水的容器

// func main() {

// }

// 双指针法，两个指针向中间移动，每次都是移动数字小的那个
func maxArea(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	l, r := 0, len(height)-1
	res := 0

	for l < r {
		if height[l] < height[r] {
			res = max(res, height[l]*(r-l))
			l++
		} else {
			res = max(res, height[r]*(r-l))
			r--
		}
	}

	return res
}

// 双指针法，两个指针向中间移动，每次都是移动数字小的那个
func maxArea1(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	res := 0
	l, r := 0, len(height)-1
	for l < r {
		if height[l] < height[r] {
			res = max(res, height[l]*(r-l))
			l++
		} else {
			res = max(res, height[r]*(r-l))
			r--
		}
	}

	return res
}
