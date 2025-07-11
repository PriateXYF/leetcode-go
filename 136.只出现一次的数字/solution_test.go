package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) (res int) {
	for _, num := range nums {
		res ^= num
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSingleNumber(t *testing.T) {

}
