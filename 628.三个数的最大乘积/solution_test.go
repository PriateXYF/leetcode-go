package leetcode

import (
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maximumProduct(nums []int) int {
	min1, min2, max1, max2, max3 := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt, math.MinInt
	for _, num := range nums {
		if num >= max1 {
			max1, max2, max3 = num, max1, max2
		} else if num >= max2 {
			max2, max3 = num, max2
		} else if num >= max3 {
			max3 = num
		}
		if num <= min1 {
			min1, min2 = num, min1
		} else if num <= min2 {
			min2 = num
		}
	}
	return max(min1*min2*max1, max1*max2*max3)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumProductOfThreeNumbers(t *testing.T) {

}
