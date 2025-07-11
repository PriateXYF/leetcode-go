package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findLUSlength(a string, b string) int {
	if len(a) != len(b) {
		return max(len(a), len(b))
	}
	if a != b {
		return len(a)
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestUncommonSubsequenceI(t *testing.T) {

}
