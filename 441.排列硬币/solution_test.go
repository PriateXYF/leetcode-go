package leetcode

import (
    "math"
    "testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func arrangeCoins(n int) int {
	return int((math.Sqrt(float64(8*n+1)) - 1) / 2)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestArrangingCoins(t *testing.T) {

}
