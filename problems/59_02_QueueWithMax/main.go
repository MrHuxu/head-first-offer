/*

题目: 定义一个队列实现函数 max 得到队列里的最大值, 要求 push, pop 和 max 函数的时间复杂度都是 O(1).

解答: 算法和上一题类似, 但是太复杂了, 留着以后看吧.

*/

package main

type queue struct {
	currentIndex   int
	data, maximums []internalData
}

type internalData struct {
	number, index int
}

func (q *queue) push(x int) {
	for len(q.maximums) != 0 && x > q.maximums[len(q.maximums)-1].number {
		q.maximums = q.maximums[0 : len(q.maximums)-1]
	}

	d := internalData{number: x, index: q.currentIndex}
	q.data = append(q.data, d)
	q.maximums = append(q.maximums, d)

	q.currentIndex++
}

func (q *queue) pop() (x int) {
	if len(q.maximums) == 0 {
		return 0
	}

	if q.maximums[0].index == q.data[0].index {
		q.maximums = q.maximums[1:len(q.maximums)]
	}

	x = q.data[0].number
	q.data = q.data[1:len(q.data)]
	return
}

func (q *queue) max() (x int) {
	if len(q.maximums) == 0 {
		return 0
	}

	return q.maximums[0].number
}

func main() {
	var q queue

	q.push(2)
	q.push(1)
	println(q.max())
	q.pop()
	println(q.max())

	q.push(4)
	println(q.max())
	q.push(3)
	println(q.max())

	println(q.pop())
	println(q.max())

	println(q.pop())
	println(q.max())

	println(q.pop())
	println(q.max())
}
