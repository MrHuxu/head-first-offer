/*

题目: 字符串转换成整数.

解答:

需要注意几个 case:

1. 如果使用 0 作为空串的返回值的话, 如何区分输入的是空串还是 "0", C++ 的 atoi 函数是用了一个全局变量标识, go 语言可以用一个 error;
2. 注意数字前面的正负号.

*/

package main

import (
	"errors"
	"fmt"
)

var err = errors.New("invalid input")

// StrToInt ...
func StrToInt(str string) (int, error) {
	if str == "" {
		return 0, err
	}

	var result int
	flag := true

	for i := 0; i < len(str); i++ {
		if str[i] <= '9' && str[i] >= '0' {
			result = result*10 + int(str[i]-'0')
		} else {
			if i == 0 {
				if str[i] == '+' || str[i] == '-' {
					flag = str[i] == '+'
				} else {
					return 0, err
				}
			} else {
				return 0, err
			}
		}
	}

	if !flag {
		result = -result
	}

	return result, nil
}

func main() {
	fmt.Println(StrToInt("123"))
	fmt.Println(StrToInt("-123"))
	fmt.Println(StrToInt(""))
	fmt.Println(StrToInt("0"))
	fmt.Println(StrToInt("0b1"))
	fmt.Println(StrToInt("01"))
}
