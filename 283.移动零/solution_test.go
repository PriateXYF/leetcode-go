package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	left, right := 0, 0
	for left < len(nums) && right < len(nums) {
		if nums[left] != 0 {
			left++
			if right <= left {
				right = left + 1
			}
		} else {
			// 找到第一个不为 0 的元素
			for ; right < len(nums) && nums[right] == 0; right++ {
			}
			if right >= len(nums) {
				break
			}
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMoveZeroes(t *testing.T) {

}
