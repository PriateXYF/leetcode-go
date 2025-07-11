package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func distributeCandies(candyType []int) (num int) {
	candyMap := map[int]bool{}
	for _, candy := range candyType {
		if _, exists := candyMap[candy]; !exists {
			candyMap[candy] = true
			num++
			if num >= len(candyType)/2 {
				return num
			}
		}
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDistributeCandies(t *testing.T) {

}
