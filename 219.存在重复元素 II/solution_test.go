package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyDuplicate(nums []int, k int) bool {
	table := map[int]int{}
	for now, num := range nums {
		// 如果存在
		if past, exists := table[num]; exists && now-past <= k {
			return true
		}
		table[num] = now
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestContainsDuplicateIi(t *testing.T) {

}
