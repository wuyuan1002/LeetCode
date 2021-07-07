package main

// 155. 最小栈

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// push(x) —— 将元素 x 推入栈中。
// pop()—— 删除栈顶的元素。
// top()—— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。

func main() {

}

// 同offer 30

// 每一次有元素入栈，都把当前的最小元素入最小栈
// 出栈时，两个栈同时弹出栈顶元素

type MinStack struct {
	data []int // 存放数据
	min  []int // 存放最小元素
}

func Constructor1() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
	}
}

// Push 压入
func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.min) == 0 || x < this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

// Pop 删除栈顶元素
func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

// Top 弹出栈顶元素
func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[len(this.data)-1]
}

// GetMin 获取当前栈中的最小值
func (this *MinStack) GetMin() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
