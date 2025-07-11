package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	sp, tp := 0, 0
	for sp < len(s) {
		if tp >= len(t) {
			return false
		}
		for t[tp] != s[sp] {
			tp++
			if tp >= len(t) {
				return false
			}
		}
		tp++
		sp++
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIsSubsequence(t *testing.T) {

}
