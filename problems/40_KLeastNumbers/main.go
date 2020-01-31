/*

题目: 输入 n 个整数, 找出其中最小的 k 个数.

解答: 我们使用一个大根堆, 遍历数字:

1. 如果堆中元素数量小于 k, 那么直接把数字压入堆;
2. 如果堆中元素数量等于 k, 则把数字和堆顶元素比较, 如果大于等于堆顶元素, 则这个数字肯定不属于最小的 k 个元素, 如果小于堆顶元素, 则弹出堆顶元素, 并把当前数字压入堆.

这样需要对时间复杂度为 O(nlogk), 而且因为这样只需要维护大小为 k 的堆, 特别适合 n 很大而 k 很小的场景, 比如找出一个超大数据流里的最小的若干个数.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/types"
)

// GetLeastNumbers ...
func GetLeastNumbers(data []int, k int) []int {
	var heap types.MaxHeap

	for _, num := range data {
		if heap.Size() < k {
			heap.Push(num)
			continue
		}

		if num < heap.Peek() {
			heap.Pop()
			heap.Push(num)
		}
	}

	var result []int
	for heap.Size() != 0 {
		result = append(result, heap.Pop())
	}
	return result
}

func main() {
	fmt.Println(GetLeastNumbers([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4))
}
