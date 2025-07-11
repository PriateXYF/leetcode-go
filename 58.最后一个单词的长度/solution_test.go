package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLastWord(s string) int {
	length := 0
	// 反向遍历
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			length++
		} else if length != 0 {
			break
		}
	}
	return length
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLengthOfLastWord(t *testing.T) {

}
