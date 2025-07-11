package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	mod := 2*numRows - 2
	lines := make([][]int32, numRows)
	for i, c := range s {
		line := i % mod
		if i%mod > mod/2 {
			line = numRows - (i%mod - mod/2) - 1
		}
		lines[line] = append(lines[line], c)
	}
	res := ""
	for _, line := range lines {
		res += string(line)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestZigzagConversion(t *testing.T) {
	convert("PAYPALISHIRING", 4)
	convert("PAYPALISHIRING", 3)
}
