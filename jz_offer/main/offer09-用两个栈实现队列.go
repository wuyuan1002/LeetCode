package main

// 用两个栈实现队列

// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// 分别完成在队列尾部插入整数和在队列头部删除整数的功能。
// (若队列中没有元素，deleteHead操作返回 -1 )

// 入队列时始终将元素放进ru栈，
// 出队列时始终从chu栈获取栈顶元素，若chu栈已经为空，将ru栈的元素挨个出栈放入chu队列，再从chu栈获取栈顶元素
func main() {
	queue := Constructor()
	queue.DeleteHead()
	queue.AppendTail(5)
	queue.AppendTail(7)
	queue.DeleteHead()
	queue.AppendTail(2)
	queue.DeleteHead()
	queue.DeleteHead()
}

type CQueue struct {
	ru  *[]int // 入队列所用的栈
	chu *[]int // 出队列所用的栈
}

func Constructor() CQueue {
	return CQueue{
		ru:  &[]int{},
		chu: &[]int{},
	}
}

// AppendTail 入队列
func (this *CQueue) AppendTail(value int) {
	*this.ru = append(*this.ru, value)
}

// DeleteHead 出队列
func (this *CQueue) DeleteHead() int {
	if len(*this.chu) == 0 && len(*this.ru) > 0 {
		for v := range *this.ru {
			*this.chu = append(*this.chu, (*this.ru)[len(*this.ru)-1-v])
		}
		*this.ru = (*this.ru)[len(*this.ru):]
	}
	val := -1
	if len(*this.chu) > 0 {
		val = (*this.chu)[len(*this.chu)-1]
		*this.chu = (*this.chu)[:len(*this.chu)-1]
	}
	return val
}

// ----------第二次做------------

// AppendTail2 入队列
func (this *CQueue) AppendTail2(value int) {
	*this.ru = append(*this.ru, value)
}

// DeleteHead2 出队列
func (this *CQueue) DeleteHead2() int {
	if len(*this.chu) == 0 && len(*this.ru) >= 0 {
		for i := len(*this.ru) - 1; i >= 0; i-- {
			*this.chu = append(*this.chu, (*this.ru)[i])
		}
		*this.ru = (*this.ru)[:0]
	}

	if len(*this.chu) > 0 {
		val := (*this.chu)[len(*this.chu)-1]
		*this.chu = (*this.chu)[:len(*this.chu)-1]
		return val
	}
	return -1
}
