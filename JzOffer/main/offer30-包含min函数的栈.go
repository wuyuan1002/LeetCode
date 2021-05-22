package main

// 剑指 Offer 30. 包含min函数的栈

// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
// 调用 min、push 及 pop 的时间复杂度都是 O(1)。

func main() {

}

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

// Min 获取当前栈中的最小值
func (this *MinStack) Min() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

// ----------------
// 第二次做

type MinStack1 struct {
	stack []int
	min   []int
}

func Constructor11() MinStack1 {
	return MinStack1{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

// Push1 压入
func (this *MinStack1) Push1(x int) {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 || x < this.min[len(this.min)-1] {
		this.stack = append(this.stack, x)
	} else {
		this.stack = append(this.stack, this.min[len(this.min)-1])
	}
}

// Pop1 删除栈顶元素
func (this *MinStack1) Pop1() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

// Top1 弹出栈顶元素
func (this *MinStack1) Top1() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

// Min1 获取当前栈中的最小值
func (this *MinStack1) Min1() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}
