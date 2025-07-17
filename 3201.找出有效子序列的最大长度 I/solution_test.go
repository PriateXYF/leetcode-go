package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maximumLength(nums []int) int {
	// 三种情况，全奇数、全偶数、奇偶相间（以数组第一个元素的奇偶性开头）
	odd, even, alternate := 0, 0, 0
	mask := 1 - nums[0]&1
	for _, num := range nums {
		odd += num & 1
		even += 1 - num&1
		alternate += (1 & num) ^ mask
		mask = ((1 & num) ^ mask) ^ mask
	}
	return max(odd, even, alternate)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindTheMaximumLengthOfValidSubsequenceI(t *testing.T) {

}
