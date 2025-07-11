package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	head, tail := 0, len(nums)-1
	for mid := (head + tail) / 2; head != tail-1 && head < tail; mid = (head + tail) / 2 {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			tail = mid
		} else {
			head = mid
		}
	}
	if len(nums) == 0 && nums[0] == target {
		return 0
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinarySearch(t *testing.T) {

}
