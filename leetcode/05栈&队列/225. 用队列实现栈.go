package main

// 225. 用队列实现栈

// 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
//
// 实现 MyStack 类：
//
// void push(int x) 将元素 x 压入栈顶。
// int pop() 移除并返回栈顶元素。
// int top() 返回栈顶元素。
// boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
//
// 注意：
//
// 你只能使用队列的基本操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
// 你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。

// MyStack .
// 同 leetcode 232. 用栈实现队列
// 使用一个队列即可模拟栈
// 入栈时元素正常添加到队列尾部
// 出栈时先将除队列尾元素外的其余元素挨个从队头移动到队尾，然后弹出元素
type MyStack struct {
	queue []int // 使用一个队列模拟栈
}

// ConstructorMyStack .
func ConstructorMyStack() MyStack {
	return MyStack{queue: []int{}}
}

// Push 入栈时元素正常添加到队列尾部
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

// Pop 出栈时先将除队列尾元素外的其余元素挨个从队头移动到队尾，然后弹出元素
func (this *MyStack) Pop() int {
	// 若队列元素大于1 -- 则将队列的前 len(this.queue)-1 个元素重新入队，移动完后新的队头元素即为原来的队尾元素
	if len(this.queue) > 1 {
		for i := 0; i < len(this.queue)-1; i++ {
			this.queue = append(this.queue, this.queue[0])
			this.queue = this.queue[1:]
		}
	}

	// 进行出栈操作 -- 获取队头元素并删除
	val := -1
	if len(this.queue) > 0 {
		val = this.queue[0]
		this.queue = this.queue[1:]
	}

	return val
}

// Top 获取栈顶元素值 -- 不进行出栈操作
func (this *MyStack) Top() int {
	val := this.Pop()
	if val != -1 {
		// 将删除的数据添加回去
		this.queue = append(this.queue, val)
	}
	return val
}

// Empty 判断栈是否为空
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
