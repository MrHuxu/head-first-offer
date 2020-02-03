/*

题目: 给定一个数组和滑动窗口的大小, 请找出所有滑动窗口里的最大值.

解答:

定义一个双端队列存储数组下标, 每次遍历到一个数字:

如果现在遍历到长度小于窗口, 就对比当前数字和前一个数字, 如果大于前一个数字, 那么当前数字可能是最大值, 存到队头, 否则不可能是, 直接丢弃.

如果数字已经遍历到窗口长度了:

1. 如果当前数字大于队头数字, 那么当前数字就是目前窗口的最大值, 压入队头;
2. 如果当前数字小于队头, 那么就从队尾弹出所有小于当前数字的值, 并把当前数字压入队尾;
3. 每次遍历到位置 i, 位置 i-size 肯定会滑出窗口, 如果滑出的是当前队头, 则弹出当前队头, 根据步骤 2, 弹出队头之后剩下的队头也是当前窗口最大值.

每次把队头的下标对应值压入结果数组返回即可.

*/

package main

import "fmt"

// maxInWindows ...
func maxInWindows(nums []int, size int) (result []int) {
	if size == 1 {
		return nums
	}

	queue := []int{0}

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if i < size-1 {
			if num > nums[i-1] {
				queue = []int{i}
			}
		} else {
			if num > nums[queue[0]] {
				queue = []int{i}
			} else {
				for i := len(queue) - 1; i >= 0; i-- {
					if num >= nums[queue[i]] {
						queue = queue[0 : len(queue)-1]
					}
				}
				queue = append(queue, i)
			}

			if queue[0] == i-size {
				queue = queue[1:len(queue)]
			}

			result = append(result, nums[queue[0]])
		}
	}

	return result
}

func main() {
	fmt.Println(maxInWindows([]int{2, 3, 4, 2, 6, 2, 5, 1}, 3))
}
