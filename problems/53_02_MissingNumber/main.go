/*

题目: 大小为 n-1 已经排序的数组, 其中缺失了 0~n-1 中的一个, 找出缺失的这一个.

解答: 二分查找, 如果当前数字等于下标, 那么缺失的数字一定在右边的一半, 否则就从左边继续找.

*/

package main

// GetMissingNumber ...
func GetMissingNumber(data []int) int {
	front := 0
	tail := len(data) - 1

	for front <= tail {
		mid := (front + tail) / 2

		if data[mid] == mid {
			front = mid + 1
		} else {
			if mid == 0 {
				return mid
			}

			if data[mid-1] == mid-1 {
				return mid
			}

			tail = mid - 1
		}
	}

	return len(data)
}

func main() {
	println(GetMissingNumber([]int{0, 1, 2, 3, 4, 5}))
	println(GetMissingNumber([]int{0, 1, 2, 3, 5, 6}))
	println(GetMissingNumber([]int{0, 1, 2, 4, 5, 6}))
}
