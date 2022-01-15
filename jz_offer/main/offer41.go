package main

// 剑指 Offer 41. 数据流中的中位数

// 如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，
// 那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，
// 那么中位数就是所有数值排序之后中间两个数的平均值。

// func main() {
// 	aa := Constructor2()
// 	aa.AddNum(1)
// 	aa.AddNum(2)
// 	// aa.AddNum(3)
// 	// aa.AddNum(4)
// 	// aa.AddNum(5)

// 	fmt.Println(aa.FindMedian())
// }

type MedianFinder1 struct {
	nums []int
}

// 1.使用未排序的数组来存储流中的数据，每次插入时直接放入数组尾部<O(1)>,
// 每次获取中位数时使用partition函数获取下标，计算中位数<O(n)>
// 也可以使用排序的数组、排序的链表、二叉搜索树、最大堆最小堆来存
func Constructor2() MedianFinder1 {
	return MedianFinder1{
		nums: make([]int, 0),
	}
}

func (this *MedianFinder1) AddNum(num int) {
	this.nums = append(this.nums, num)
}

func (this *MedianFinder1) FindMedian() float64 {
	numsLen := len(this.nums)
	l, r := 0, 0
	if numsLen%2 == 0 {
		l = (numsLen - 1) / 2
		r = (numsLen + 1) / 2
	} else {
		l = numsLen / 2
		r = l
	}

	var media float64
	findLeft := false
	findRight := false
	left, right := 0, numsLen-1
	index := partition1(this.nums, left, right)
	for {
		if index == l {
			findLeft = true
			if l == r { // 奇数个
				media = float64(this.nums[l])
				break
			}
			if findRight {
				media = (media + float64(this.nums[l])) / 2
				break
			}
			media = float64(this.nums[l])
			left = r
			// index = partition1(this.nums, r, numsLen-1)
		} else if index == r {
			findRight = true
			if l == r { // 奇数个
				media = float64(this.nums[l])
				break
			}
			if findLeft {
				media = (media + float64(this.nums[r])) / 2
				break
			}
			media = float64(this.nums[r])
			right = r - 1
			// index = partition1(this.nums, 0, r-1)
		} else if index < l {
			left = index + 1
		} else if index > r {
			right = index - 1
		}
		index = partition1(this.nums, left, right)
	}
	return media
}

func partition1(arr []int, left, right int) int {

	l, r := left, right
	temp := arr[left]
	for l < r {
		for arr[r] >= temp && l < r {
			r--
		}
		for arr[l] <= temp && l < r {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[left], arr[l] = arr[l], temp
	return l
}
