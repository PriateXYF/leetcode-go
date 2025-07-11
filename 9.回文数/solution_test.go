package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 判断整数是否为回文
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	revertNum := 0
	for revertNum < x {
		revertNum = revertNum*10 + x%10
		x = x / 10
	}
	return revertNum == x || (revertNum/10 == x)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPalindromeNumber(t *testing.T) {

}
