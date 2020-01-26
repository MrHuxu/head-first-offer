/*

题目: 用两个栈实现一个队列.

解答: 每次入队的时候都向 stack1 压栈, 出队的时候, 分两种情况:

1. 如果 stack2 是空的, 那么把 stack1 里的元素全部出栈并且压入 stack2 中, 这样 stack2 的栈顶就是最先压入 stack1 的元素, 直接从 stack2 出栈即可;
2. 如果 stack2 不是空的, 那么根据1, stack2 的栈顶就是最先入队元素, 从 stack2 出栈即可.

使用栈的一个好处就是, 栈每次压栈出栈就改变了元素的顺序, 可以做很灵活的操作, 使用队列模拟栈的话, 为了拿到队列的最后一个元素来模拟栈顶, 就需要把一个队列完全转移, 其实性能很差.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/types"
)

type queue struct {
	stack1 types.Stack
	stack2 types.Stack
}

func (q *queue) enqueue(x interface{}) {
	q.stack1.Push(x)
}

func (q *queue) dequeue() (x interface{}) {
	if q.stack2.Size() != 0 {
		x = q.stack2.Pop()
		return
	}

	if q.stack1.Size() == 0 {
		return
	}

	for q.stack1.Size() != 0 {
		q.stack2.Push(q.stack1.Pop())
	}
	x = q.stack2.Pop()
	return
}

func main() {
	var q queue

	q.enqueue(1)
	q.enqueue(2)
	fmt.Println(q.dequeue())
	q.enqueue(4)
	fmt.Println(q.dequeue())
	q.enqueue(3)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	q.enqueue(5)
	fmt.Println(q.dequeue())
}
