package leetcode

import (
	"testing"
)

func isBadVersion(version int) bool {
	if version < 1 {
		return false
	}
	return true
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 查找版本出错的第一个错误版本
func firstBadVersion(n int) (mid int) {
	if n == 1 {
		return n
	}
	if isBadVersion(1) {
		return 1
	}
	head, tail := 1, n
	for mid = (head + tail) / 2; isBadVersion(mid) == isBadVersion(mid+1); mid = (head + tail) / 2 {
		if isBadVersion(mid) {
			tail = mid
		} else {
			head = mid
		}
	}
	return mid + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFirstBadVersion(t *testing.T) {

}
