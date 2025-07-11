package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isSelfDividingNumber(num int) bool {
	for temp := num; temp != 0; temp = temp / 10 {
		carry := temp % 10
		if carry == 0 || num%carry != 0 {
			return false
		}
	}
	return true
}

// 判断范围内的数是否是自除数
func selfDividingNumbers(left int, right int) (ans []int) {
	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			ans = append(ans, i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSelfDividingNumbers(t *testing.T) {

}
