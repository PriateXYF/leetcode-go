package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findSpecialInteger(arr []int) int {
	combo := 1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			combo++
			if combo > len(arr)/4 {
				return arr[i]
			}
		} else {
			combo = 1
		}
	}
	return arr[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestElementAppearingMoreThan25InSortedArray(t *testing.T) {

}
