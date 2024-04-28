package main

// 剑指 Offer 59 - II. 队列的最大值

// 请定义一个队列并实现函数 max_value 得到队列里的最大值，
// 要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
//
// 若队列为空，pop_front 和 max_value需要返回 -1

// func main() {
// 	obj := Constructor4()
// 	// obj.Pop_front()
// 	// obj.Pop_front()
// 	// obj.Pop_front()
// 	// obj.Pop_front()
// 	// obj.Pop_front()
// 	// obj.Push_back(15)
// 	// fmt.Println(obj.Max_value())
// 	// obj.Push_back(9)
// 	// fmt.Println(obj.Max_value())

// 	obj.Push_back(1)
// 	obj.Push_back(2)
// 	fmt.Println(obj.Max_value())
// 	obj.Pop_front()
// 	fmt.Println(obj.Max_value())
// }

type MaxQueue struct {
	queue    []int // 存储队列中的值 -- 先进先出
	maxQueue []int // 辅助队列，单调递减 -- 第一个数是队列的最大值，第二个数是辅助队列第一个数后面队列中数据的最大值，第三个数是辅助队列第二个数后面队列中数据的最大值，以此类推
}

func Constructor4() MaxQueue {
	return MaxQueue{
		queue:    make([]int, 0),
		maxQueue: make([]int, 0),
	}
}

// Max_value 获取最大值
func (this *MaxQueue) Max_value() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.maxQueue[0]
}

// Push_back 插入元素到尾部
func (this *MaxQueue) Push_back(value int) {

	this.queue = append(this.queue, value)

	index := 0 // 第一个小于value的下标 -- 这里之后的都删掉

	// 如果不为空并且插入的值小于最大值，则从maxQueue中把小于value的数删除掉，把value插入到队列末尾；否则把队列清空，之后直接把value添加到队列中
	if len(this.maxQueue) > 0 && value < this.maxQueue[0] {
		for index = len(this.maxQueue) - 1; index >= 0; index-- {
			if this.maxQueue[index] >= value {
				index++
				break
			}
		}
	}

	// 将index之后的元素全部删掉，把value插入到队列末尾
	this.maxQueue = append(this.maxQueue[:index], value)
}

// Pop_front 从队头删掉一个元素
func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}

	value := this.queue[0]
	this.queue = this.queue[1:] // 删掉队列的第一个元素
	if value == this.maxQueue[0] {
		this.maxQueue = this.maxQueue[1:] // 若要删掉的元素是最大值，则同时删掉最大值
	}

	return value
}
