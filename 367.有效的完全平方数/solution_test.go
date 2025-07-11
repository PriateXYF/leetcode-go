package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	left, right := 1, num
	var mid int
	for mid = (left + right) / 2; !((mid*mid) <= num && (mid+1)*(mid+1) > num); mid = (left + right) / 2 {
		if (mid * mid) > num {
			right = mid
		} else if (mid+1)*(mid+1) <= num {
			left = mid
		}
	}
	return mid*mid == num
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidPerfectSquare(t *testing.T) {

}
