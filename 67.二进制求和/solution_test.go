package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 两个二进制数相加，返回当前位和是否进位
func add(a uint8, b uint8, carry uint8) (rune, uint8) {
	if a+b+carry == 3 {
		return '1', 1
	} else if a+b+carry == 2 {
		return '0', 1
	} else if a+b+carry == 1 {
		return '1', 0
	} else if a+b+carry == 0 {
		return '0', 0
	}
	return 'x', 0
}

// 对二进制字符串进行求和操作
func addBinary(a string, b string) string {
	long, short := a, b
	if len(a) < len(b) {
		long, short = short, long
	}
	var res []rune
	// 进位标识符
	var carry uint8 = 0
	// 遍历长字符串
	for li := len(long) - 1; li >= 0; li-- {
		si := len(short) - (len(long) - li)
		var sc, lc uint8 = 0, long[li] - '0'
		if si >= 0 {
			sc = short[si] - '0'
		}
		var el rune
		el, carry = add(sc, lc, carry)
		res = append(res, el)
	}
	if carry == 1 {
		res = append(res, '1')
	}
	// 逆转结果
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAddBinary(t *testing.T) {

}
