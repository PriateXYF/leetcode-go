package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) (res int) {
	res = nums[0]
	quantity := 1
	for i := 1; i < len(nums); i++ {
		if res == nums[i] {
			quantity++
		} else {
			if quantity == 0 {
				res = nums[i]
			} else {
				quantity--
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMajorityElement(t *testing.T) {

}
