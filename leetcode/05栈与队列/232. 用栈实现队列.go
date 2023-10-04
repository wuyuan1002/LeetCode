package main

// 232. 用栈实现队列

// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//
// 实现 MyQueue 类：
//
// void push(int x) 将元素 x 推到队列的末尾
// int pop() 从队列的开头移除并返回元素
// int peek() 返回队列开头的元素
// boolean empty() 如果队列为空，返回 true ；否则，返回 false
// 说明：
//
// 你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
// 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。

// MyQueue .
// 同 Offer 09. 用两个栈实现队列
// 入队列时始终将元素放进ru栈，
// 出队列时始终从chu栈获取栈顶元素，若chu栈已经为空，将ru栈的元素挨个出栈放入chu栈，再从chu栈获取栈顶元素
type MyQueue struct {
	ru  []int // 入队列所用的栈
	chu []int // 出队列所用的栈
}

// Constructor .
func Constructor() MyQueue {
	return MyQueue{
		ru:  []int{},
		chu: []int{},
	}
}

// Push 入队列时始终将元素放进ru栈
func (this *MyQueue) Push(x int) {
	this.ru = append(this.ru, x)
}

// Pop 出队列时始终从chu栈获取栈顶元素，若chu栈已经为空，将ru栈的元素挨个出栈放入chu栈，再从chu栈获取栈顶元素
func (this *MyQueue) Pop() int {
	// 若chu栈为空了且入栈有元素，则将ru栈元素出栈添加到chu栈中
	if len(this.chu) == 0 && len(this.ru) > 0 {
		for i := len(this.ru) - 1; i >= 0; i-- {
			this.chu = append(this.chu, this.ru[i])
		}
		this.ru = this.ru[len(this.ru):]
	}

	// 进行出队列操作
	val := -1
	if len(this.chu) > 0 {
		val = this.chu[len(this.chu)-1]
		this.chu = this.chu[:len(this.chu)-1]
	}
	return val
}

// Peek 仅获取队列头元素的值 -- 不出队列
func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val != -1 {
		// 将删除的数据添加回去
		this.chu = append(this.chu, val)
	}
	return val
}

// Empty 队列是否为空
func (this *MyQueue) Empty() bool {
	return len(this.ru) == 0 && len(this.chu) == 0
}
