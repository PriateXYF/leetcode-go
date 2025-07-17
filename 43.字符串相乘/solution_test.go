package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 字符串一位乘法
func multiplyOne(num1 string, num2 int32) (res string) {
	if num2 == '0' {
		return "0"
	}
	var carry, remainder int
	for i := len(num1) - 1; i >= 0; i-- {
		num := num1[i]
		a, b := int(num-'0'), int(num2-'0')
		ans := a*b + carry
		carry, remainder = ans/10, ans%10
		res = strconv.Itoa(remainder) + res
	}
	if carry != 0 {
		res = strconv.Itoa(carry) + res
	}
	return res
}

// 字符串加法
func add(num1 string, num2 string) (res string) {
	var carry, remainder int
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var a, b int
		if i < 0 {
			a = 0
		} else {
			a = int(num1[i] - '0')
		}
		if j < 0 {
			b = 0
		} else {
			b = int(num2[j] - '0')
		}
		ans := a + b + carry
		carry, remainder = ans/10, ans%10
		res = strconv.Itoa(remainder) + res
	}
	if carry != 0 {
		res = strconv.Itoa(carry) + res
	}
	return
}
func multiply(num1 string, num2 string) (res string) {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	for _, base := range num1 {
		// 进行一位乘法
		temp := multiplyOne(num2, base)
		res = add(res, temp) + "0"
	}
	return res[:len(res)-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMultiplyStrings(t *testing.T) {
	res := multiply("1234", "1230")
	fmt.Println(res)
}
