package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	table := map[int]int{}
	for i, num := range nums {
		div := target - nums[i]
		if j, exists := table[div]; exists {
			return []int{j, i}
		} else {
			table[num] = i
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTwoSum(t *testing.T) {

}
