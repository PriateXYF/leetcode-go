package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseString(t *testing.T) {

}
