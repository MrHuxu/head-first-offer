/*

题目: 我们把只包含因子 2, 3 和 5 的数称作丑数, 求按从小到大的第 n 个丑数.

解答:

假设已经有一些丑数, 其中最大值是 m, 那么存在一个数字 n2, 使 n2*2>m, 且比 n2 小的数字乘以二都小于 m, 同理可以找到这样的 n3 和 n5 两个数字, 那么下一个丑数肯定是 min(n2*2, n3*3, n5*5), 我们找到下一个数字后, 再更新这三个数字, 最后找到第 n 个即可.

*/

package main

// GetUglyNumber ...
func GetUglyNumber(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 3
	}
	if n == 3 {
		return 5
	}

	nums := make([]int, n)

	nums[0] = 2
	nums[1] = 3
	nums[2] = 5
	n2 := 1
	n3 := 0
	n5 := 0
	for i := 3; i < n; i++ {
		nums[i] = min(nums[n2]*2, nums[n3]*3, nums[n5]*5)

		for ; n2 < n; n2++ {
			if nums[n2]*2 > nums[i] {
				break
			}
		}
		for ; n3 < n; n3++ {
			if nums[n3]*3 > nums[i] {
				break
			}
		}
		for ; n5 < n; n5++ {
			if nums[n5]*5 > nums[i] {
				break
			}
		}
	}

	return nums[n-1]
}

func min(a, b, c int) int {
	min := a
	if b < a {
		min = b
	}
	if c < min {
		min = c
	}

	return min
}

func main() {
	println(GetUglyNumber(1))
	println(GetUglyNumber(4))
	println(GetUglyNumber(1500))
}
