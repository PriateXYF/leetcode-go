package leetcode

import (
	"testing"
)

func guess(num int) int {
	res := 2
	if num == res {
		return 0
	} else if num > res {
		return -1
	} else {
		return 1
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

// 猜数字大小
func guessNumber(n int) int {
	left, right := 1, n
	var mid int
	if guess(left) == 0 {
		return left
	} else if guess(right) == 0 {
		return right
	}
	for mid = (left + right) / 2; guess(mid) != 0; mid = (left + right) / 2 {
		if guess(mid) == -1 {
			right = mid
		} else if guess(mid) == 1 {
			left = mid
		}
	}
	return mid
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGuessNumberHigherOrLower(t *testing.T) {

}
