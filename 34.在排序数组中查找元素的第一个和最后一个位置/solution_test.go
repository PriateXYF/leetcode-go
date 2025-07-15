package leetcode

import (
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	first := sort.SearchInts(nums, target)
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}
	second := sort.SearchInts(nums, target+1)
	return []int{first, second - 1}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindFirstAndLastPositionOfElementInSortedArray(t *testing.T) {

}
