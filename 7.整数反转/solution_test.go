package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) (res int) {
	for ; x != 0; x = x / 10 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		res = res*10 + x%10
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseInteger(t *testing.T) {
	x := int32(-1 << 31)
	fmt.Println(x, x-1)
	//reverse(-3214)
}
