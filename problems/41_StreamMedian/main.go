/*

题目: 数据流中的中位数.

解答:

建立一个大根堆一个小根堆, 使用大根堆存储较小的一半数字, 小根堆存储较大的一半数字, 我们轮流向这两个堆压入数字.

向大根堆压数字时:

1. 如果小于等于堆顶元素, 则直接压入堆;
2. 如果大于堆顶元素, 说明这个元素属于大的那一半, 则把这个数字压入小根堆, 再从小根堆弹出一个数字压入大根堆.

向小根堆压数字时:

1. 如果大于等于堆顶元素, 则直接压入堆;
2. 如果小于堆顶元素, 说明这个元素属于小的那一半, 则把这个数字压入大根堆, 再从大根堆弹出一个数字压入小根堆.

最后如果两个堆数字数量和为偶数, 则中位数为大根堆和小根堆堆顶数字和除以 2. 如果为奇数, 因为先开始压的是大根堆, 大根堆元素会多一个, 直接弹出大根堆堆顶数字即可.


*/

package main

import "github.com/MrHuxu/types"

// GetMedian ...
func GetMedian(numbers []int) float64 {
	var smallerHeap types.MaxHeap
	var largerHeap types.MinHeap

	for i, number := range numbers {
		if i%2 == 0 {
			if smallerHeap.Size() == 0 {
				smallerHeap.Push(number)
			} else {
				if number <= smallerHeap.Peek() {
					smallerHeap.Push(number)
				} else {
					largerHeap.Push(number)
					smallerHeap.Push(largerHeap.Pop())
				}
			}
		} else {
			if largerHeap.Size() == 0 {
				largerHeap.Push(number)
			} else {
				if number >= largerHeap.Peek() {
					largerHeap.Push(number)
				} else {
					smallerHeap.Push(number)
					largerHeap.Push(smallerHeap.Pop())
				}
			}
		}
	}

	if smallerHeap.Size() == largerHeap.Size() {
		return (float64(smallerHeap.Pop()) + float64(largerHeap.Pop())) / 2.0
	}
	return float64(smallerHeap.Pop())
}

func main() {
	println(GetMedian([]int{4, 5, 1, 6, 2, 7, 3}))
	println(GetMedian([]int{4, 5, 1, 6, 2, 7, 3, 8}))
}
