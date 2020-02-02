/*

题目: 数字在排序数组中出现的次数.

解答: 通过二分法找数字出现的边界, 而不是简单遍历.

*/

package main

// GetNumberOfK ...
func GetNumberOfK(data []int, k int) int {
	return getLastK(data, k) - getFirstK(data, k) + 1
}

func getFirstK(data []int, k int) int {
	front := 0
	tail := len(data) - 1

	for front <= tail {
		mid := (front + tail) / 2

		if data[mid] == k {
			if mid == 0 {
				return mid
			} else if data[mid-1] != data[mid] {
				return mid
			} else {
				tail = mid - 1
			}
		} else if data[mid] > k {
			tail = mid - 1
		} else {
			front = mid + 1
		}
	}

	return 0
}

func getLastK(data []int, k int) int {
	front := 0
	tail := len(data) - 1

	for front <= tail {
		mid := (front + tail) / 2

		if data[mid] == k {
			if mid == len(data)-1 {
				return mid
			} else if data[mid+1] != data[mid] {
				return mid
			} else {
				front = mid + 1
			}
		} else if data[mid] > k {
			tail = mid - 1
		} else {
			front = mid + 1
		}
	}

	return 0
}

func main() {
	println(GetNumberOfK([]int{1, 2, 3, 3, 3, 3, 4, 5}, 3))
}
