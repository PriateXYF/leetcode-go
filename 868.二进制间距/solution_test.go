package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func binaryGap(n int) int {
	start, dist, res := false, 0, 0
	for ; n != 0; n = n >> 1 {
		if n&1 == 1 {
			start = true
			res = max(dist, res)
			dist = 1
		} else if start {
			dist++
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryGap(t *testing.T) {
	binaryGap(8)
}
