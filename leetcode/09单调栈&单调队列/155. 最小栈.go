package main

// 155. 最小栈

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 实现 MinStack 类:
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。

// MinStack .
// 同 Offer 59-2、leetcode 239. 滑动窗口最大值
type MinStack struct {
	stack    []int // 存储栈中所有元素 -- 后进先出
	minstack []int // 单调递减栈 -- 栈顶元素为最小值，栈顶的下一个元素为栈内栈顶元素后的最小值，以此类推
}

// ConstructorMinStack .
func ConstructorMinStack() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minstack: make([]int, 0),
	}
}

// Push 向栈内添加元素 -- 若添加的元素小于等于单调栈的栈顶元素，则同时入栈到单调栈中，即更新栈内的最小值
func (this *MinStack) Push(val int) {
	// 进行入栈操作
	this.stack = append(this.stack, val)

	// 单调栈为空或添加的元素小于等于当前栈的最小值，则将新元素加入到单调栈中，更新栈的最小值
	if len(this.minstack) == 0 || val <= this.minstack[len(this.minstack)-1] {
		this.minstack = append(this.minstack, val)
	}
}

// Pop 弹出栈顶元素 -- 若弹出的元素正好是最小值，则同时从单调栈中弹出，即更新栈内最小值
func (this *MinStack) Pop() {
	// 若要弹出的元素正好是最小值，则同时从单调栈中弹出，更新栈的最小值
	if len(this.minstack) > 0 && this.stack[len(this.stack)-1] == this.minstack[len(this.minstack)-1] {
		this.minstack = this.minstack[:len(this.minstack)-1]
	}

	// 进行出栈操作
	this.stack = this.stack[:len(this.stack)-1]
}

// Top 获取栈顶元素
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// GetMin 获取栈内最小值
func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}
