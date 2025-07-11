package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func intToRoman(num int) (res string) {
	table := map[int][]string{
		0: {"", "", "", ""},
		1: {"I", "X", "C", "M"},
		2: {"II", "XX", "CC", "MM"},
		3: {"III", "XXX", "CCC", "MMM"},
		4: {"IV", "XL", "CD"},
		5: {"V", "L", "D"},
		6: {"VI", "LX", "DC"},
		7: {"VII", "LXX", "DCC"},
		8: {"VIII", "LXXX", "DCCC"},
		9: {"IX", "XC", "CM"},
	}
	for i := 0; num != 0; i++ {
		digit := num % 10
		res = table[digit][i] + res
		num /= 10
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIntegerToRoman(t *testing.T) {
	intToRoman(3749)
	intToRoman(58)
	intToRoman(1994)
}
