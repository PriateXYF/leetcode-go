package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findPoisonedDuration(timeSeries []int, duration int) (res int) {
	start, end := 0, 0
	for _, time := range timeSeries {
		if time > end {
			res += end - start + 1
			start = time
			end = time + duration - 1
		} else if time <= end {
			end = time + duration - 1
		}
	}
	res += end - start
	if timeSeries[0] == 0 {
		res++
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTeemoAttacking(t *testing.T) {

}
