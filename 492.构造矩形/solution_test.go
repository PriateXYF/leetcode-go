package leetcode

import (
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for area%w > 0 {
		w--
	}
	return []int{area / w, w}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConstructTheRectangle(t *testing.T) {

}
