package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	head := 1
	tail := x
	var sqrt int
	// 只有 sqrt^2 <= x 且 (sqrt+1)^2 > x 才满足条件
	for sqrt = (head + tail) / 2; !(sqrt*sqrt <= x && (sqrt+1)*(sqrt+1) > x); sqrt = (head + tail) / 2 {
		// 如果 sqrt^2 大于 x ，则目标值偏大
		if sqrt*sqrt > x {
			tail = sqrt
		} else if (sqrt+1)*(sqrt+1) <= x { // 如果 (sqrt+1)^2 小于 x ，则目标值偏小
			head = sqrt
		}
	}
	return sqrt
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSqrtx(t *testing.T) {

}
