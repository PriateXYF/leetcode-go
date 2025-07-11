package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func countSegments(s string) (res int) {
	for i := 0; i < len(s)-1; i++ {
		if s[i] != ' ' && s[i+1] == ' ' {
			res++
		}
	}
	if len(s) != 0 && s[len(s)-1] != ' ' {
		res++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNumberOfSegmentsInAString(t *testing.T) {

}
