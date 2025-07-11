package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[0] > target || letters[len(letters)-1] <= target {
		return letters[0]
	}
	left, right := 0, len(letters)-1
	var mid int
	// 二分查找先找到目标字母
	for mid = left + (right-left)/2; left <= right; mid = left + (right-left)/2 {
		if mid+1 < len(letters) && letters[mid] < target && letters[mid+1] > target {
			return letters[mid+1]
		}
		if letters[mid] == target {
			break
		}
		if letters[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	for i := mid; i < len(letters); i++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindSmallestLetterGreaterThanTarget(t *testing.T) {

}
