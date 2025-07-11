package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func containsDuplicate(nums []int) bool {
	table := map[int]bool{}
	for _, num := range nums {
		if _, exists := table[num]; exists {
			return true
		}
		table[num] = true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestContainsDuplicate(t *testing.T) {

}
