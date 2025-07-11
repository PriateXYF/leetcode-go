package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 查找一个数的因数
func findFactor(n int) (factor int) {
	factors := []int{5, 3, 2}
	for _, f := range factors {
		if n%f == 0 {
			return f
		}
	}
	return -1
}

// 判断一个整数是否为丑数
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		if factor := findFactor(n); factor == -1 {
			return false
		} else {
			n /= factor
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestUglyNumber(t *testing.T) {

}
