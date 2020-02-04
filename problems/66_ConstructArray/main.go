/*

题目: 给定一个数组 A[0,1,...,n-1], 构建一个数组 B[0,1,...,n-1], 其中 B 中元素 B[i]=A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1], 不能使用除法.

解答: 每一行分别计算的话, 时间复杂度是 O(n^2), 这样太慢了, 可以把 B[i] 分成 C[i]=A[0]*A[1]*...*A[i-1] 和 D[i]=A[i+1]*...*A[n-1] 两部分, 可以从前到后得到 C[i]=C[i-1]*A[i-1], 从后向前可以得到 D[i]=D[i+1]*A[i+1], 最后两者相乘得到 B[i], 时间复杂度为 O(n).

*/

package main

import "fmt"

// multiply ...
func multiply(arr []int) []int {
	result := make([]int, len(arr))

	c := make([]int, len(arr))
	c[0] = 1
	for i := 1; i < len(c); i++ {
		c[i] = c[i-1] * arr[i-1]
	}

	d := make([]int, len(arr))
	d[len(arr)-1] = 1
	for i := len(d) - 2; i >= 0; i-- {
		d[i] = d[i+1] * arr[i+1]
	}

	for i := range result {
		result[i] = c[i] * d[i]
	}

	return result
}

func main() {
	fmt.Println(multiply([]int{1, 2, 3, 4, 5, 6, 7}))
}
