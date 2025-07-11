package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func replaceElements(arr []int) []int {
	maxVal := arr[len(arr)-1]
	// 逆序遍历
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= maxVal {
			arr[i], maxVal = maxVal, arr[i]
		} else {
			arr[i] = maxVal
		}
	}
	arr[len(arr)-1] = -1
	return arr
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReplaceElementsWithGreatestElementOnRightSide(t *testing.T) {

}
