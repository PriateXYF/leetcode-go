package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isHappy(n int) bool {
	// 重复字典
	repeatDict := map[int]bool{}
	for n != 1 {
		// 如果该数重复
		if _, exists := repeatDict[n]; exists {
			return false
		}
		repeatDict[n] = true
		num := n
		happy := 0
		for num > 0 {
			happy += (num % 10) * (num % 10)
			num /= 10
		}
		n = happy
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestHappyNumber(t *testing.T) {

}
