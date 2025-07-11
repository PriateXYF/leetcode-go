package leetcode

import (
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func arrayPairSum(nums []int) (res int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestArrayPartition(t *testing.T) {

}
