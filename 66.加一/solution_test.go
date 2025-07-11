package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func plusOne(digits []int) []int {
	// 是否需要扩展一位
	extFlag := true
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			extFlag = false
			break
		}
	}
	if extFlag {
		digits = append([]int{1}, digits...)
	}
	return digits
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPlusOne(t *testing.T) {

}
