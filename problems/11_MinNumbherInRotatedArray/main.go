/*

题目: 输入一个递增排序的数组的一个旋转, 输出旋转数组的最小元素.

解答:

我们可以采用类似二分的思路来解, 首先声明 left 和 right 指针为数组的起始和结束位置, 那么 left 一定在左侧的递增数组, right 一定在右侧的递增数组

对两者的中间位置 mid 进行判断:

1. 如果 numbers[mid] >= numbers[left], 那么 mid 一定位于左侧的递增数组, 那么我们可以把 left 移动到 mid;
2. 如果 numbers[mid] <= numbers[left], 那么 mid 一定位于右侧的递增数组, 那么我们可以把 right 移动到 mid.

移动到最后, left 和 right 一定在相邻位置, 此时 numbers[right] 就是最小的值了.

*/

package main

// Min ...
func Min(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for numbers[left] >= numbers[right] {
		if right-left == 1 {
			return numbers[right]
		}

		mid := (left + right) / 2

		if numbers[mid] >= numbers[left] {
			left = mid
		} else if numbers[mid] <= numbers[left] {
			right = mid
		} else {
			return numbers[mid]
		}
	}

	return numbers[right]
}

func main() {
	println(Min([]int{3, 4, 5, 6, 7, 2}))
}
