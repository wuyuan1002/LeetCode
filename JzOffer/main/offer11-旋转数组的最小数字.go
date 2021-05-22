package main

import "fmt"

// 旋转数组的最小数字

// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
// 例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
func main() {
	fmt.Printf("%d", minArray([]int{1, 2, 3}))
	// fmt.Printf("%d", minArray([]int{2, 3, 4, 0, 1, 2}))
	// fmt.Printf("%d", minArray([]int{10, 1, 10, 10, 10}))
	// fmt.Printf("%d", minArray([]int{10, 10, 1, 10, 10, 10, 10}))
}

// 使用二分法查找
func minArray(numbers []int) int {
	if numbers == nil || len(numbers) == 0 {
		panic("数组为nil或长度为0")
	}
	p := 0                 // 第一个元素下标
	q := len(numbers) - 1  // 最后一个元素下标
	mid := p               // p和q中间的元素下标
	if len(numbers) == 1 { // 若数组中只有一个元素
		return numbers[p]
	}
	for numbers[p] >= numbers[q] {
		if p+1 == q {
			mid = q
			break
		}
		mid = (p + q) / 2
		if numbers[mid] == numbers[p] && numbers[mid] == numbers[q] {
			// 若三个指向的值相等，则无法判断中间的数字是位于前面的数组还是后面的数组，只能顺序查找
			return orderFind(numbers[p : q+1])
		}
		if numbers[mid] >= numbers[p] {
			p = mid
		} else if numbers[mid] <= numbers[q] {
			q = mid
		}
	}
	return numbers[mid]
}

func orderFind(arr []int) int {
	min := 0 // 已找到最小值的下标
	for i := range arr {
		if arr[i] < arr[min] {
			min = i
		}
	}
	return arr[min]
}

// 第二次做
func minArray1(numbers []int) int {
	if numbers == nil || len(numbers) == 0 {
		panic("数组为nil或长度为0")
	}
	if len(numbers) == 1 {
		return numbers[0]
	}

	first := 0               // 第一个元素下标
	last := len(numbers) - 1 // 最后一个元素下标
	mid := first             // 中间元素的下标; 把中间下标初始化为first，如果是把数组的前0个元素移到数组后面，其实就递增数组本身，那么最小值其实就是数组的第一个元素，在后面的循环中就不会进入，这也是为什么循环条件不能写成(first < last)的原因

	// first始终往大的数字上移动，last始终往小的数字上移动，当两个下标的差为1时，说明first移到了最大的数字上，last移到了最小的数字上
	for numbers[first] >= numbers[last] { // 循环条件不可以是first < last，见mid的解释
		if last-first == 1 {
			mid = last
			break
		}
		mid = (first + last) / 2

		if numbers[mid] == numbers[first] && numbers[mid] == numbers[last] {
			// 若三个指向的值相等，则无法判断中间的数字是位于前面的数组还是后面的数组，只能顺序查找
			return orderFind(numbers[first : last+1])
		}

		if numbers[mid] >= numbers[first] {
			first = mid
		} else if numbers[mid] <= numbers[last] {
			last = mid
		}
	}
	return numbers[mid]
}
