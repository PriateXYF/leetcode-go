package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	str := []rune(s)
	// 计算总共需要反转的次数
	carry := len(s) % (k * 2)
	times := (len(s) - carry) / (k * 2)
	if carry >= k {
		times += 1
	}
	for i := 0; i < times; i++ {
		start, end := i*k*2, i*k*2+k-1
		for start < end {
			str[start], str[end] = str[end], str[start]
			start++
			end--
		}
	}
	if carry != 0 && carry < k {
		start, end := len(s)-carry, len(s)-1
		for start < end {
			str[start], str[end] = str[end], str[start]
			start++
			end--
		}
	}
	return string(str)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseStringIi(t *testing.T) {

}
