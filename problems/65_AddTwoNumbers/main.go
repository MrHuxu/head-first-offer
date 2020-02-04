/*

题目: 不使用四则运算实现加法.

解答: 使用位操作求解.

1. 第一步将各个位相加, 不计算进位, 相加得到的数字和异或的结果相同(1+0==1^0, 0+0==0^0, (1+1)>>1==1^1);
2. 第二步计算进位, 需要进位的数字和按位与然后左移一位的结果相同(原理同上);
3. 将上面两步的结果相加, 直到进位为 0 为止.

*/

package main

// Add ...
func Add(num1, num2 int) int {
	sum := num1 ^ num2
	plus := (num1 & num2) << 1

	for plus != 0 {
		num1 = sum
		num2 = plus

		sum = num1 ^ num2
		plus = (num1 & num2) << 1
	}

	return sum
}

func main() {
	println(Add(1, 2))
	println(Add(12, 59))
}
