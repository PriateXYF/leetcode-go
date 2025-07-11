package leetcode

import (
	"strconv"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) (res string) {
	// 保证 num1 为较长字符串
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}
	// 补充前导零
	for len(num1) != len(num2) {
		num2 = "0" + num2
	}
	carry := false
	for i := len(num1) - 1; i >= 0; i-- {
		digit := int(num1[i]-'0') + int(num2[i]-'0')
		if carry {
			digit++
		}
		if digit >= 10 {
			digit = digit - 10
			carry = true
		} else {
			carry = false
		}
		res = strconv.Itoa(digit) + res
	}
	if carry {
		res = "1" + res
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAddStrings(t *testing.T) {

}
