/*

题目: 在数组中的两个数字, 如果前面一个数字大于后面的数字, 则这两个数字组成一个逆序对. 输入一个数组, 求出这个数组里逆序对的总数.

解答:

使用递归的思想, 每次把数组分成两部分, 把两个数组合并成有序的, 并且统计前面数组里比后面大的个数就是逆序对的个数.

因为两个子数组在处理完之后都是从小到大有序的, 那么我们从后向前遍历:

1. 假设第一个数组中第 i1 个数字大于第二个数组中第 i2 个数字, 那么 i1 肯定比 i2 之前的数字都大, 直接把结果累加上 i2, 向前移动 i1;
2. 假设第一个数组中第 i1 个数字小于等于第二个数组中第 i2 个数字, 肯定不是一个逆序对, 向前移动 i2.

最后返回总对数即可.

*/

package main

// InversePairs ...
func InversePairs(data []int) int {
	_, cnt := inversePairsHelper(data)
	return cnt
}

func inversePairsHelper(data []int) ([]int, int) {
	if len(data) == 1 {
		return data, 0
	}

	result := make([]int, len(data))
	sub1, cnt1 := inversePairsHelper(data[0 : len(data)/2])
	sub2, cnt2 := inversePairsHelper(data[len(data)/2 : len(data)])

	cnt := cnt1 + cnt2

	i1 := len(sub1)
	i2 := len(sub2)
	for {
		if i1 == 0 && i2 == 0 {
			break
		}

		switch {
		case i1 == 0:
			result[i1+i2-1] = sub2[i2-1]
			i2--

		case i2 == 0:
			result[i1+i2-1] = sub1[i1-1]
			i1--

		default:
			if sub1[i1-1] > sub2[i2-1] {
				cnt += i2

				result[i1+i2-1] = sub1[i1-1]
				i1--
			} else {
				result[i1+i2-1] = sub2[i2-1]
				i2--
			}
		}
	}

	return result, cnt
}

func main() {
	println(InversePairs([]int{7, 5, 6, 4}))
}
