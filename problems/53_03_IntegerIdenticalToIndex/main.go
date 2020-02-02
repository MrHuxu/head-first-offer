/*

题目: 一个递增排序的数组, 找出数组中和下标相等的元素.

解答: 仍然使用二分解法, 如果找到数字和下标一样, 直接返回, 否则数字大于下标, 那在左侧继续寻找, 否则从右侧查找.

*/

package main

// GetNumberSameAsIndex ...
func GetNumberSameAsIndex(numbers []int) int {
	front := 0
	tail := len(numbers) - 1

	for front < tail {
		mid := (front + tail) / 2

		if numbers[mid] == mid {
			return mid
		} else if numbers[mid] > mid {
			tail = mid - 1
		} else {
			front = mid + 1
		}
	}

	return numbers[front]
}

func main() {
	println(GetNumberSameAsIndex([]int{-3, -1, 1, 3, 5}))
}
