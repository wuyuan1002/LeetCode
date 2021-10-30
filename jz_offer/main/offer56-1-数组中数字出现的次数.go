package main

// 剑指 Offer 56 - I. 数组中数字出现的次数

// 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
// 请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

func main() {

}

// 136，260
func singleNumbers(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	flag := 0 // 找到分组依据 -- 把数组中数字异或后，相同数字都抵消为0了，相当于是两个不同数字进行了异或，所以，flag的二进制位为1的就是两个数字不同的位
	for _, n := range nums {
		flag ^= n
	}

	mask := flag & -flag // 找到flag中第一个为1的位 -- 即找到两个数字的第一个不相同的位

	result := make([]int, 2)
	for _, n := range nums {
		if n&mask == 0 {
			result[0] ^= n
		} else {
			result[1] ^= n
		}
	}
	return result
}

// 第二次做
func singleNumbers1(nums []int) []int {
	flag := 0
	for _, num := range nums {
		flag ^= num
	}

	mask := flag & -flag // 找到flag中第一个为1的位

	result := make([]int, 2)
	for _, num := range nums {
		if num&mask == 0 {
			result[0] ^= num
		} else {
			result[1] ^= num
		}
	}
	return result
}
