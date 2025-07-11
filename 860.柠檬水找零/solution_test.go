package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func lemonadeChange(bills []int) bool {
	wallet := map[int]int{5: 0, 10: 0}
	for _, bill := range bills {
		switch bill {
		case 5:
			wallet[5]++
		case 10:
			if wallet[5] <= 0 {
				return false
			}
			wallet[10]++
			wallet[5]--
		case 20:
			if wallet[10] <= 0 {
				if wallet[5] < 3 {
					return false
				}
				wallet[5] -= 3
			} else if wallet[5] < 1 {
				return false
			} else {
				wallet[10]--
				wallet[5]--
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLemonadeChange(t *testing.T) {

}
