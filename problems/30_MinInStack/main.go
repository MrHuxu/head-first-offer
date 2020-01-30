/*

题目: 定义栈的数据结构, 请在该类型中实现一个能得到最小元素的 min 函数, push, pop 和 min 函数的时间复杂度为 O(1).

解答: 可以使用一个辅助栈来保存当前栈里的最小值, 每次压栈的同时也把当前最小值压到这个辅助栈里, 弹栈的同时也对这个辅助栈弹栈, 通过这个辅助栈就可以 O(1) 来获取最小值.

*/

package main

type stack struct {
	data1 []int
	data2 []int
}

func (s *stack) push(v int) {
	s.data1 = append(s.data1, v)

	if len(s.data2) == 0 {
		s.data2 = append(s.data2, v)
	} else {
		if v < s.data2[len(s.data2)-1] {
			s.data2 = append(s.data2, v)
		} else {
			s.data2 = append(s.data2, s.data2[len(s.data2)-1])
		}
	}
}

func (s *stack) pop() (v int) {
	if len(s.data1) == 0 {
		return -1
	}

	v = s.data1[len(s.data1)-1]
	s.data1 = s.data1[0 : len(s.data1)-1]
	s.data2 = s.data2[0 : len(s.data2)-1]
	return
}

func (s *stack) min() int {
	if len(s.data2) == 0 {
		return -1
	}

	return s.data2[len(s.data2)-1]
}

func main() {
	var s stack
	s.push(3)
	println(s.min())
	s.push(4)
	println(s.min())
	s.push(2)
	s.push(1)
	println(s.pop())
	println(s.pop())
	s.push(0)
	println(s.min())
}
