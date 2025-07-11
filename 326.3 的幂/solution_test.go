package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfThree(n int) bool {
	for ; n > 0 && n%3 == 0; n /= 3 {
	}
	return n == 1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPowerOfThree(t *testing.T) {

}
