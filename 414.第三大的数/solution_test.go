package leetcode

import (
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func thirdMax(nums []int) int {
	first, second, third := math.MinInt, math.MinInt, math.MinInt
	for _, num := range nums {
		if num > first {
			first, second, third = num, first, second
		} else if num < first && num > second {
			second, third = num, second
		} else if num < second && num > third {
			third = num
		}
	}
	if third == math.MinInt {
		return first
	}
	return third
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThirdMaximumNumber(t *testing.T) {

}
