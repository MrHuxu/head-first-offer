/*

题目: 输入两个整数序列, 第一个序列表示栈顶压入顺序, 清判断第二个序列是否为该栈顶弹出顺序.

解答:

使用一个栈来模拟压栈弹栈过程, 遍历 pop 数组, 如果当前栈为空或栈顶元素不为当前遍历到的元素, 则将 push 中元素取出压栈直到栈顶元素为当前元素为止, 并且弹栈, 再重复这个过程.

如果 push 数组提前被取光, 那么不是弹出顺序.

如果 stack 最后为空, 那么是弹出顺序.

*/

package main

// IsPopOrder ...
func IsPopOrder(push, pop []int) bool {
	var stack []int

	for _, num := range pop {
		for len(stack) == 0 || stack[len(stack)-1] != num {
			if len(push) == 0 {
				return false
			}

			stack = append(stack, push[0])
			push = push[1:len(push)]
		}

		stack = stack[0 : len(stack)-1]
	}

	return len(stack) == 0
}

func main() {
	println(IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	println(IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 2, 1}))
	println(IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
