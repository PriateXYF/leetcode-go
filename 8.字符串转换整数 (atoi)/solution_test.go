package leetcode

import (
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func myAtoi(s string) int {
	res, start, negative := 0, false, false
	for _, c := range s {
		if !start {
			// 第一个字符非法
			if c != ' ' && c != '+' && c != '-' && (c < '0' || c > '9') {
				return res
			} else if c == ' ' {
				continue
			} else if c == '-' {
				start, negative = true, true
				continue
			} else if c == '+' {
				start, negative = true, false
				continue
			} else {
				start = true
			}
		}
		// 进入处理阶段
		if c < '0' || c > '9' {
			break
		} else {
			num := int(c) - '0'
			if negative {
				num = -num
			}
			res = res*10 + num
		}
		if res > math.MaxInt32 {
			res = math.MaxInt32
			break
		} else if res < math.MinInt32 {
			res = math.MinInt32
			break
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestStringToIntegerAtoi(t *testing.T) {

}
