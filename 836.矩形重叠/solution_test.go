package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	overlap := rec1[3] == rec2[1] || rec1[1] == rec2[3] || rec1[0] == rec2[2] || rec1[2] == rec2[0]
	hover := rec1[3] < rec2[1] || rec1[1] > rec2[3] || rec1[0] > rec2[2] || rec1[2] < rec2[0]
	return !hover && !overlap
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRectangleOverlap(t *testing.T) {

}
