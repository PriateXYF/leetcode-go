package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func missingNumber(nums []int) (res int) {
	for i, num := range nums {
		res += i
		res -= num
	}
	res += len(nums)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMissingNumber(t *testing.T) {

}
